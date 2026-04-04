package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Murat")
	want := "Merhaba, Murat"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
