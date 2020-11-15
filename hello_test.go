package main

import "testing"

func TestHello(t *testing.T) { /* t is our hook to the testing environment*/

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper() // tells the test suite that this method is a helper. (line numbers)
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Mike", "English")
		want := "Hello, Mike!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to the world", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "¡Hola, Elodie!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Anna", "French")
		want := "Bonjour, Anna!"
		assertCorrectMessage(t, got, want)
	})
}
