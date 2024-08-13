package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(got, want, t)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(got, want, t)
	})

	t.Run("saying hello in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(got, want, t)
	})

	t.Run("saying hello in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"
		assertCorrectMessage(got, want, t)
	})

	t.Run("saying hello in Hindi", func(t *testing.T) {
		got := Hello("Elodie", "Hindi")
		want := "नमस्ते, Elodie"
		assertCorrectMessage(got, want, t)
	})
}

func assertCorrectMessage(got, want string, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
