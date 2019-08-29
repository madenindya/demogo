package main

import (
	"testing"
	// "fmt"
)

func TestSum(t *testing.T) {
	expected := 10
	a, b := 5, 5

	c := Sum(a, b)
	if c != expected {
		t.Errorf("Not true")
	}
}
