package main

import (
	"bytes"
	"net/http"
	"os"
	"testing"
)

func TestGreet(t *testing.T) {

	t.Run("print to a Buffer", func(t *testing.T) {

		buffer := bytes.Buffer{}
		Greet(&buffer, "Chris")

		got := buffer.String()
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("print to Stdout", func(t *testing.T) {

		Greet(os.Stdout, "Elodie") // os.Stdout and bytes.Buffer implements Writer

	})

	t.Run("print to Internet", func(t *testing.T) {
		http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
		Greet(os.Stdout, "Elodie") // os.Stdout and bytes.Buffer implements Writer

	})
}
