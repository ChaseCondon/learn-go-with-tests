package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chase")
	want := "Hello, Chase!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
