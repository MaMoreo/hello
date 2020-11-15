package main

import "testing"

func TestHello(t *testing.T) { /* t is our hook to the testing environment*/
	got := Hello("Mike")
	want := "Hello, Mike!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
