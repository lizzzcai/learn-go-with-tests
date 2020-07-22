package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Lize")

	got := buffer.String()
	want := "Hello, Lize"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
