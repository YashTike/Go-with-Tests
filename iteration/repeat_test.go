package iteration

import (
	"fmt"
	"testing"
)

func TestRepeater(t *testing.T) {

	t.Run("repeating 5 times", func(t *testing.T) {

		iter := Repeat("a", 5)
		expected := "aaaaa"

		assertCorrectOutput(iter, expected, t)
	})

	t.Run("repeating 1 time", func(t *testing.T) {

		iter := Repeat("a", 1)
		expected := "a"

		assertCorrectOutput(iter, expected, t)
	})
}

func assertCorrectOutput(got, want string, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}
