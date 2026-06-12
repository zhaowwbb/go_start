package main

import (
	"reflect"
	"testing"
)

func TestChristmasLights(t *testing.T) {
	cl := NewChristmasLights(8)

	expectedSteps := [][]bool{
		{true, true, false, false, false, false, false, true}, // Step 1: pos moves to 1
		{true, false, true, false, false, false, false, true}, // Step 2: pos moves to 2
		{true, false, false, true, false, false, false, true}, // Step 3: pos moves to 3
		{true, false, false, false, true, false, false, true}, // Step 4: pos moves to 4
		{true, false, false, false, false, true, false, true}, // Step 5: pos moves to 5
		{true, false, false, false, false, false, true, true}, // Step 6: pos moves to 6 (reaches boundary, flips direction)
		{true, false, false, false, false, true, false, true}, // Step 7: pos moves back to 5
		{true, false, false, false, true, false, false, true}, // Step 8: pos moves back to 4
		{true, false, false, true, false, false, false, true}, // Step 9: pos moves back to 3
		{true, false, true, false, false, false, false, true}, // Step 10: pos moves back to 2
		{true, true, false, false, false, false, false, true}, // Step 11: pos moves back to 1 (reaches boundary, flips direction)
	}

	for i, expected := range expectedSteps {
		stepNum := i + 1
		actual := cl.Next()
		if !reflect.DeepEqual(actual, expected) {
			// fmt.Println()
			t.Errorf("Step %d failed:\n expected %v, actual %v", stepNum, expected, actual)
		}
	}
}
