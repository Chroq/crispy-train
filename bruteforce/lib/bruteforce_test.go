package lib_test

import (
	"bruteforce/lib"
	"testing"
)

/* func TestShen(t *testing.T) {
	expected := "Sh3n"
	got := lib.FindPasswordNaive("Sh3n")

	if got != expected {
		t.Errorf("Expected %v but got %v", expected, got)
	}
} */

func BenchmarkAkaliNaive(b *testing.B) {
	for b.Loop() {
		lib.FindPasswordNaive("Ak4l1")
	}
}

func BenchmarkAkaliCombination(b *testing.B) {
	for b.Loop() {
		lib.FindPasswordByCombination("Ak4l1")
	}
}

func BenchmarkAkaliAsync(b *testing.B) {
	for b.Loop() {
		lib.FindPasswordAsync("Ak4l1")
	}
}

func BenchmarkAkaliCombinationAsync(b *testing.B) {
	for b.Loop() {
		lib.FindPasswordByCombinationAsync("Ak4l1")
	}
}
