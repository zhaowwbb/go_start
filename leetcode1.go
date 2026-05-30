package main

// twoSum finds the indices of the two numbers that add up to the target
func twoSum(nums []int, target int) []int {
	// Create a map to store the number as the key and its index as the value
	numToIndex := make(map[int]int)

	// Iterate through the array
	for i, num := range nums {
		complement := target - num

		// If the complement exists in our map, we found the pair
		if index, exists := numToIndex[complement]; exists {
			return []int{index, i}
		}

		// Otherwise, store the current number and its index
		numToIndex[num] = i
	}

	// Return an empty slice if no solution is found
	return []int{}
}

func twoSumV2(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if index, exists := indexMap[complement]; exists {
			return []int{index, i}
		} else {
			indexMap[num] = i
		}
	}
	return []int{}
}

func twoSumV3(nums []int, target int) []int {
	numToIndex := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if index, exists := numToIndex[complement]; exists {
			return []int{index, i}
		}
		numToIndex[num] = i
	}

	return []int{}
}
