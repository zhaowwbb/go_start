package main

// RemoveElement removes all occurrences of val in nums in-place.
// It returns the number of elements which are not equal to val (k).
func RemoveElement(nums []int, val int) int {
	// writeIndex tracks where the next non-val element should be placed
	writeIndex := 0

	// Loop through all elements in the slice
	for readIndex := 0; readIndex < len(nums); readIndex++ {
		// If the current element is not equal to the target value
		if nums[readIndex] != val {
			nums[writeIndex] = nums[readIndex]
			writeIndex++
		}
	}

	return writeIndex
}

func RemoveElementV2(nums []int, val int) int {
	writeIndex := 0

	for readIndex := 0; readIndex < len(nums); readIndex++ {
		if nums[readIndex] != val {
			nums[writeIndex] = nums[readIndex]
			writeIndex++
		}
	}
	return writeIndex
}

func RemoveElementV3(nums []int, val int) int {
	writeIndex := 0
	for readIndex := 0; readIndex < len(nums); readIndex++ {
		if nums[readIndex] != val {
			nums[writeIndex] = nums[readIndex]
			writeIndex++
		}
	}

	return writeIndex
}
