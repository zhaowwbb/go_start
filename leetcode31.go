package main

// NextPermutation modifies nums in-place to the next lexicographical permutation.
func NextPermutation(nums []int) {
	// 1. Find first decreasing index from the right
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	// If no such index, reverse entire slice
	if i == -1 {
		reverse(nums, 0, len(nums)-1)
		return
	}

	// 2. Find element just larger than nums[i]
	j := len(nums) - 1
	for nums[j] <= nums[i] {
		j--
	}

	// 3. Swap
	nums[i], nums[j] = nums[j], nums[i]

	// 4. Reverse suffix
	reverse(nums, i+1, len(nums)-1)
}

// helper: reverse slice in-place
func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func NextPermutationV2(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i == -1 {
		reverse(nums, 0, len(nums)-1)
		return
	}
	j := len(nums) - 1
	for nums[i] >= nums[j] {
		j--
	}

	nums[i], nums[j] = nums[j], nums[i]

	reverse(nums, i+1, len(nums)-1)
}

func reverseV2(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
