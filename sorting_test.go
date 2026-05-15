package main

import (
	"fmt"
	"reflect"
	"testing"
)

// TestAllSortingImplementations executes validation against all three QuickSort flavors
func TestAllSortingImplementations(t *testing.T) {
	// Define structural test cases
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Standard unsorted array",
			input:    []int{10, 7, 8, 9, 1, 5},
			expected: []int{1, 5, 7, 8, 9, 10},
		},
		{
			name:     "Already sorted array",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Array with duplicate values",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5},
			expected: []int{1, 1, 2, 3, 4, 5, 5, 6, 9},
		},
		{
			name:     "Empty slice safety catch",
			input:    []int{},
			expected: []int{},
		},
	}

	// Matrix list of the sorting function references to sweep through
	funcsToTest := []struct {
		label string
		fn    func([]int)
	}{
		{"Sort_V1", Sort},
		// {"Sort_V2", Sort2},
		// {"Sort_V3", Sort3},
		{"Sort_V4", Sort4},
		{"Sort_V5", Sort5},
	}

	// Loop through each function variation, then run every test case scenario against it
	for _, target := range funcsToTest {
		t.Run(target.label, func(t *testing.T) {
			fmt.Println(target.label)
			for _, tc := range tests {
				t.Run(tc.name, func(t *testing.T) {
					// Duplicate input slice to avoid side-effect mutations across assertions

					fmt.Println("Input:", tc.input)
					testInput := make([]int, len(tc.input))

					copy(testInput, tc.input)

					// Execute the sorting mutation
					target.fn(testInput)
					fmt.Println("Output:", testInput)

					// Deep check equality verification between sliced collections
					if !reflect.DeepEqual(testInput, tc.expected) {
						t.Errorf("[%s - %s] failed: got %v, want %v", target.label, tc.name, testInput, tc.expected)
					}
				})
			}
			fmt.Println("================================")
			// )
		})
	}
}
