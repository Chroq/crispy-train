package lib_test

import (
	"bruteforce/lib"
	"testing"
)

func TestShen(t *testing.T) {
	expected := "Sh3n"
	got := lib.FindPasswordNaive("Sh3n")

	if got != expected {
		t.Errorf("Expected %v but got %v", expected, got)
	}
}

func BenchmarkZedNaive(b *testing.B) {
	for b.Loop() {
		lib.FindPasswordNaive("z3d")
	}
}

func BenchmarkZedCombination(b *testing.B) {
	for b.Loop() {
		lib.FindPasswordByCombination("z3d")
	}
}
