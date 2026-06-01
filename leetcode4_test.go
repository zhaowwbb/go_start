package main

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected float64
	}{
		{"Example 1: Odd combined total", []int{1, 3}, []int{2}, 2.0},
		{"Example 2: Even combined total", []int{1, 2}, []int{3, 4}, 2.5},
		{"First array empty", []int{}, []int{1}, 1.0},
		{"Second array empty", []int{3}, []int{}, 3.0},
		{"Overlapping arrays", []int{1, 2}, []int{1, 2, 3}, 2.0},
		{"Includes negative integers", []int{-5, 3, 6, 12, 15}, []int{-12, -2, 0, 11, 20}, 4.5},
		{"Completely disjoint sets", []int{1, 2, 3}, []int{4, 5, 6, 7, 8}, 4.5},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// actual := FindMedianSortedArrays(tc.nums1, tc.nums2)
			// actual := FindMedianSortedArraysV2(tc.nums1, tc.nums2)
			actual := FindMedianSortedArraysV3(tc.nums1, tc.nums2)
			if actual != tc.expected {
				t.Errorf("FindMedianSortedArrays(%v, %v) = %f; expected %f", tc.nums1, tc.nums2, actual, tc.expected)
			}
		})
	}
}
