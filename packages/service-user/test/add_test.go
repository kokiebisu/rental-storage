package test

import (
	"testing"
)

func add(a int, b int) int {
	return a + b
}

// for a valid return value.
func TestAdd(t *testing.T) {
	sum := add(1, 2)
	if sum != 3 {
		t.Fatalf("Expected result is different")
	}
}
