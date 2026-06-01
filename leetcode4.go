// Package leetcode4 contains the solution for LeetCode 4: Median of Two Sorted Arrays.
package main

import (
	"math"
)

// FindMedianSortedArrays finds the median of two sorted arrays in O(log(min(m,n))) time.
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Ensure nums1 is the shorter array to optimize binary search range
	if len(nums1) > len(nums2) {
		return FindMedianSortedArrays(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	low, high := 0, m
	totalLeft := (m + n + 1) / 2

	for low <= high {
		// Calculate partition cuts
		i := (low + high) / 2
		j := totalLeft - i

		// Assign boundary values. Handle edge partitions using safely casted bounds.
		var nums1LeftMax, nums1RightMin int
		if i == 0 {
			nums1LeftMax = math.MinInt32
		} else {
			nums1LeftMax = nums1[i-1]
		}

		if i == m {
			nums1RightMin = math.MaxInt32
		} else {
			nums1RightMin = nums1[i]
		}

		var nums2LeftMax, nums2RightMin int
		if j == 0 {
			nums2LeftMax = math.MinInt32
		} else {
			nums2LeftMax = nums2[j-1]
		}

		if j == n {
			nums2RightMin = math.MaxInt32
		} else {
			nums2RightMin = nums2[j]
		}

		// Check if partition is valid
		if nums1LeftMax <= nums2RightMin && nums2LeftMax <= nums1RightMin {
			// Odd combined total elements
			if (m+n)%2 != 0 {
				return float64(max(nums1LeftMax, nums2LeftMax))
			}
			// Even combined total elements
			return float64(max(nums1LeftMax, nums2LeftMax)+min(nums1RightMin, nums2RightMin)) / 2.0
		} else if nums1LeftMax > nums2RightMin {
			// Partition cut is too far right in nums1; slide left
			high = i - 1
		} else {
			// Partition cut is too far left in nums1; slide right
			low = i + 1
		}
	}

	return 0.0
}

// Helper local functions to manage int clean boundaries (built into Go 1.21+)
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func FindMedianSortedArraysV2(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return FindMedianSortedArraysV2(nums2, nums1)
	}
	m := len(nums1)
	n := len(nums2)
	totalLeft := (m + n + 1) / 2
	left, right := 0, m
	for left <= right {
		p1 := (left + right) / 2
		p2 := totalLeft - p1

		var leftMax1, rightMin1 int
		if p1 == 0 {
			leftMax1 = math.MinInt32
		} else {
			leftMax1 = nums1[p1-1]
		}
		if p1 == m {
			rightMin1 = math.MaxInt32
		} else {
			rightMin1 = nums1[p1]
		}

		var leftMax2, rightMin2 int
		if p2 == 0 {
			leftMax2 = math.MinInt32
		} else {
			leftMax2 = nums2[p2-1]
		}
		if p2 == n {
			rightMin2 = math.MaxInt32
		} else {
			rightMin2 = nums2[p2]
		}

		if leftMax1 <= rightMin2 && leftMax2 <= rightMin1 {
			if (m+n)%2 != 0 {
				return float64(max(leftMax1, leftMax2))
			} else {
				return float64(max(leftMax1, leftMax2)+min(rightMin1, rightMin2)) / 2.0
			}
		} else if leftMax1 > rightMin2 {
			right = p1 - 1
		} else {
			left = p1 + 1
		}
	}

	return 0.0
}

func FindMedianSortedArraysV3(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return FindMedianSortedArraysV3(nums2, nums1)
	}
	m, n := len(nums1), len(nums2)
	low, high := 0, m
	totalLeft := (m + n + 1) / 2

	for low <= high {
		p1 := (low + high) / 2
		p2 := totalLeft - p1

		var leftMax1, rightMin1 int
		if p1 == 0 {
			leftMax1 = math.MinInt32
		} else {
			leftMax1 = nums1[p1-1]
		}
		if p1 == m {
			rightMin1 = math.MaxInt32
		} else {
			rightMin1 = nums1[p1]
		}
		var leftMax2, rightMin2 int
		if p2 == 0 {
			leftMax2 = math.MinInt32
		} else {
			leftMax2 = nums2[p2-1]
		}
		if p2 == n {
			rightMin2 = math.MaxInt32
		} else {
			rightMin2 = nums2[p2]
		}

		if leftMax1 <= rightMin2 && leftMax2 <= rightMin1 {
			if (m+n)%2 != 0 {
				return float64(max(leftMax1, leftMax2))
			} else {
				return float64(max(leftMax1, leftMax2)+min(rightMin1, rightMin2)) / 2.0
			}
		} else if leftMax1 > rightMin2 {
			high = p1 - 1
		} else {
			low = p1 + 1
		}

	}
	return 0.0
}
