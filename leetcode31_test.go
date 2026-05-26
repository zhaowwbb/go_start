package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 1, 5}, []int{1, 5, 1}},
		{[]int{1, 3, 2}, []int{2, 1, 3}},
		{[]int{2, 3, 1}, []int{3, 1, 2}},
		{[]int{5, 4, 7, 5, 3, 2}, []int{5, 5, 2, 3, 4, 7}},
		{[]int{4, 2, 0, 2, 3, 2, 0}, []int{4, 2, 0, 3, 0, 2, 2}},
	}

	for _, tc := range testCases {
		// Copy input so original is preserved
		nums := append([]int{}, tc.input...)

		// NextPermutation(nums)
		NextPermutationV2(nums)

		pass := reflect.DeepEqual(nums, tc.expected)
		fmt.Printf("Input: %v Expected: %v Actual: %v Result: %v\n",
			tc.input, tc.expected, nums, passStatus(pass))

		if !reflect.DeepEqual(nums, tc.expected) {
			t.Errorf("Input: %v | Expected: %v | Actual: %v",
				tc.input, tc.expected, nums)
		}
	}
}

func passStatus(ok bool) string {
	if ok {
		return "PASS"
	}
	return "FAIL"
}
