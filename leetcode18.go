package main

import "sort"

// FourSum finds all unique quadruplets in the array which give the sum of target.
// Time Complexity: O(N^3) | Space Complexity: O(1) (ignoring output storage allocation)
func FourSum(nums []int, target int) [][]int {
	var results [][]int
	sort.Ints(nums)
	n := len(nums)

	for i := 0; i < n-3; i++ {
		// Skip duplicates for the first pointer
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Pruning 1: The lowest possible sum is already larger than target
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		// Pruning 2: The largest possible sum with this 'i' is smaller than target
		if nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}

		for j := i + 1; j < n-2; j++ {
			// Skip duplicates for the second pointer
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			// Pruning 3: Lowest possible sum here is too large
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}
			// Pruning 4: Largest possible sum here is too small
			if nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}

			// Classical Two-Pointer setup for the final two elements
			left := j + 1
			right := n - 1

			for left < right {
				currentSum := nums[i] + nums[j] + nums[left] + nums[right]

				if currentSum == target {
					results = append(results, []int{nums[i], nums[j], nums[left], nums[right]})

					// Bypass duplicate values for left and right elements
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}

					left++
					right--
				} else if currentSum < target {
					left++
				} else {
					right--
				}
			}
		}
	}

	return results
}

func FourSumV2(nums []int, target int) [][]int {
	var results [][]int
	if len(nums) < 4 {
		return results
	}
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-3; i++ {
		//remove duplicate
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			//remove duplicate
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left := j + 1
			right := n - 1
			for left < right {
				total := nums[i] + nums[j] + nums[left] + nums[right]
				if total == target {
					results = append(results, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if total < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return results
}
