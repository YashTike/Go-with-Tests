package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(slices ...[]int) (total []int) {

	for _, slyc := range slices {
		total = append(total, Sum(slyc))
	}
	return total
}

func SumAllTails(slices ...[]int) (tailTotal []int) {
	for _, slyc := range slices {
		if len(slyc) == 0 {
			tailTotal = append(tailTotal, 0)
			continue
		}
		tailTotal = append(tailTotal, Sum(slyc)-slyc[0])
	}
	return tailTotal
}
