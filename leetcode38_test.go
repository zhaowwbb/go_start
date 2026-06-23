package main

import (
	"testing"
)

func TestCountAndSay(t *testing.T) {
	// Multi-case datasets
	testInputs := []int{1, 2, 3, 4, 5}
	expectedOutputs := []string{"1", "11", "21", "1211", "111221"}

	t.Log("--- Running Count and Say Tests ---")

	// Single function call execution within the loop
	for i := 0; i < len(testInputs); i++ {
		currentInput := testInputs[i]
		expected := expectedOutputs[i]

		// Single location where the function is executed
		// actual := CountAndSay(currentInput)
		// actual := CountAndSayV2(currentInput)
		actual := CountAndSayV3(currentInput)

		// Validation check
		if actual == expected {
			t.Logf("Test Case %d: PASSED (%d -> \"%s\")", i+1, currentInput, actual)
		} else {
			t.Errorf("Test Case %d: FAILED! Input: %d | Expected: \"%s\", but got: \"%s\"",
				i+1, currentInput, expected, actual)
		}
	}
}
