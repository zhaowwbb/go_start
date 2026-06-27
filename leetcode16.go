package main

import (
	"math"
	"sort"
)

// ThreeSumClosest finds three integers in nums such that the sum is closest to target.
// Time Complexity: O(N^2) | Space Complexity: O(1) (depending on sort implementation)
func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closestSum := math.MaxInt32

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {
			currentSum := nums[i] + nums[left] + nums[right]

			if currentSum == target {
				return currentSum
			}

			// Update closestSum if currentSum is closer to the target
			if abs(currentSum-target) < abs(closestSum-target) {
				closestSum = currentSum
			}

			if currentSum < target {
				left++
			} else {
				right--
			}
		}
	}

	return closestSum
}

// Helper function to handle absolute values for integers
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ThreeSumClosestV2(nums []int, target int) int {
	sort.Ints(nums)
	closeSum := math.MaxInt32
	for i := 0; i < len(nums)-2; i++ {
		//skip duplicate
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1
		for left < right {
			currentSum := nums[i] + nums[left] + nums[right]
			if currentSum == 0 {
				return 0
			} else {
				if abs(currentSum-target) < abs(closeSum-target) {
					closeSum = currentSum
				}

				if currentSum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return closeSum
}

func ThreeSumClosestV3(nums []int, target int) int {
	sort.Ints(nums)
	closeSum := math.MaxInt32
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1
		for left < right {
			currentSum := nums[i] + nums[left] + nums[right]
			if currentSum == target {
				return 0
			} else {
				if abs(currentSum-target) < abs(closeSum-target) {
					closeSum = currentSum
				}
				if currentSum < target {
					left++
				} else {
					right--
				}
			}
		}
	}

	return closeSum
}
