package main

// searchRange implements the O(log n) dual-binary-search to locate the target boundaries.
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	firstPos := findBound(nums, target, true)
	if firstPos == -1 {
		return []int{-1, -1}
	}

	lastPos := findBound(nums, target, false)
	return []int{firstPos, lastPos}
}

func searchRangeV2(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	firstPos := findBoundV2(nums, target, true)
	if firstPos == -1 {
		return []int{-1, -1}
	}

	lastPos := findBoundV2(nums, target, false)
	return []int{firstPos, lastPos}
}

func searchRangeV3(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	firstPos := findBoundV3(nums, target, true)
	if firstPos == -1 {
		return []int{-1, -1}
	}

	secondPos := findBoundV3(nums, target, false)

	return []int{firstPos, secondPos}
}

// findBound handles searching for either the leftmost or rightmost boundary based on the flag.
func findBound(nums []int, target int, isFirst bool) int {
	left := 0
	right := len(nums) - 1
	bound := -1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			bound = mid
			if isFirst {
				right = mid - 1 // Narrow down to look for smaller indices on the left
			} else {
				left = mid + 1 // Narrow down to look for larger indices on the right
			}
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return bound
}

func findBoundV2(nums []int, target int, isFirst bool) int {
	left, right := 0, len(nums)-1
	bound := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			bound = mid
			if isFirst {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return bound
}

func findBoundV3(nums []int, target int, isFirst bool) int {
	left, right := 0, len(nums)-1
	bound := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			bound = mid
			if isFirst {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return bound
}
