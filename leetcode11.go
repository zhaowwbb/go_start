package main

// MaxArea calculates the maximum amount of water a container can store.
// Time Complexity: O(n), Space Complexity: O(1)
func MaxArea(height []int) int {
	maxWater := 0
	left := 0
	right := len(height) - 1

	for left < right {
		width := right - left

		// Find the limiting height
		currentHeight := height[left]
		if height[right] < currentHeight {
			currentHeight = height[right]
		}

		// Calculate current capacity
		currentArea := width * currentHeight
		if currentArea > maxWater {
			maxWater = currentArea
		}

		// Move the pointer that points to the shorter line inward
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxWater
}

func MaxAreaV2(height []int) int {
	maxArea := 0
	left := 0
	right := len(height) - 1
	for left < right {
		width := right - left

		currentHeight := height[left]
		if height[right] < height[left] {
			currentHeight = height[right]
		}
		currentArea := width * currentHeight
		if currentArea > maxArea {
			maxArea = currentArea
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}
