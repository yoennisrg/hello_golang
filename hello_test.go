package hello

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello, world."
	if got := Hello(); got != expected {
		t.Errorf("Expected %v got %v", expected, got)
	}
}

func TestProverb(t *testing.T) {
	expected := "Concurrency is not parallelism."
	if got := Proverb(); got != expected {
		t.Errorf("Proverb() = %q, want %q", got, expected)
	}
}
