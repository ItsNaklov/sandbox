package main

import (
	"testing"
)

func TestWalk(t *testing.T) {

	expected := "Chris"
	var got []string

	x := struct {
		Name string
	}{expected}
	walk(x, func(input string) {

	})

	if len(got) != 1 {
		t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
	}
}
