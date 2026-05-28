package main

// SearchInsert finds the target index or the insertion index in O(log n) time.
func SearchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		// Prevents potential integer overflow
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid // Target found
		} else if nums[mid] < target {
			left = mid + 1 // Target is in the right half
		} else {
			right = mid - 1 // Target is in the left half
		}
	}

	// If not found, 'left' holds the correct insertion index
	return left
}

func SearchInsertV2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
