package main

// RemoveDuplicates removes duplicates from a sorted slice in-place.
// It returns the number of unique elements (k).
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// writeIndex tracks where the next unique element should be placed
	writeIndex := 1

	// Iterate through the slice starting from the second element
	for readIndex := 1; readIndex < len(nums); readIndex++ {
		// If the current element is different from the previous one, it's unique
		if nums[readIndex] != nums[readIndex-1] {
			nums[writeIndex] = nums[readIndex]
			writeIndex++
		}
	}

	return writeIndex
}

func RemoveDuplicatesV2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	writeIndex := 1

	for readIndex := 1; readIndex < len(nums); readIndex++ {
		if nums[readIndex] != nums[readIndex-1] {
			nums[writeIndex] = nums[readIndex]
			writeIndex++
		}
	}
	return writeIndex
}

func RemoveDuplicatesV3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	writeIndex := 1

	for readIndex := 1; readIndex < len(nums); readIndex++ {
		if nums[readIndex] != nums[readIndex-1] {
			nums[writeIndex] = nums[readIndex]
			writeIndex++
		}
	}
	return writeIndex
}
