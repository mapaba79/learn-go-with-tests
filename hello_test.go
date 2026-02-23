package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Alice")
	want := "Hello, Alice"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
