package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, want, got string) {
		/*
			t.Helper() is needed to tell the test suite that this method is a helper.
			By doing this when it fails the line number reported will be in our function call rather than inside our test helper.
		*/
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", want, got)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Lize", "")
		want := "Hello, Lize"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in Spanish", func(t *testing.T) {
		got := Hello("Lize", "Spanish")
		want := "Hola, Lize"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in French", func(t *testing.T) {
		got := Hello("Lize", "French")
		want := "Bonjour, Lize"
		assertCorrectMessage(t, got, want)
	})

}
