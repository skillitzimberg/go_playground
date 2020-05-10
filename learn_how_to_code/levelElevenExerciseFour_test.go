package main

import (
	"testing"
)

func TestLevelElevenExerciseFour(t *testing.T) {
	sr, err := sqrt(-1)
	if sr != 4 {
		t.Error("Expected 4, but got . Error is .\n", sr, err)
	}

	if err != nil {
		t.Error("Expected no errors, but got .\n", err)
	}
}
