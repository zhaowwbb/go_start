package main

// search implements the O(log n) binary search on a rotated sorted array.
func searchInRotatedSortedArray(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		// Check if the left half is normally sorted
		if nums[left] <= nums[mid] {
			// Check if target lies within the sorted left half range
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1 // Search left
			} else {
				left = mid + 1 // Search right
			}
		} else { // Otherwise, the right half must be normally sorted
			// Check if target lies within the sorted right half range
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1 // Search right
			} else {
				right = mid - 1 // Search left
			}
		}
	}

	return -1
}

func searchInRotatedSortedArrayV2(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func searchInRotatedSortedArrayV3(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
