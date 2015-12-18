package main

import (
	"testing"
)

func Test_ribbonLength(t *testing.T) {
	if ribbonLength(2, 3, 4) != 34 {
		t.Errorf("Expected 2x3x4 to equal 34 feet, got %d", ribbonLength(2, 3, 4))
	}
}
