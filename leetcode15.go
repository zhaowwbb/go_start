package main

import "sort"

// ThreeSum finds all unique triplets in the array which gives the sum of zero.
func ThreeSum(nums []int) [][]int {
	var results [][]int
	sort.Ints(nums)
	length := len(nums)

	for i := 0; i < length-2; i++ {
		// If the current lowest number is greater than 0,
		// three positive numbers cannot sum to 0.
		if nums[i] > 0 {
			break
		}

		// Skip duplicate values for the first element
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := length - 1

		for left < right {
			total := nums[i] + nums[left] + nums[right]

			if total < 0 {
				left++
			} else if total > 0 {
				right--
			} else {
				results = append(results, []int{nums[i], nums[left], nums[right]})

				// Skip duplicate values for the second element
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// Skip duplicate values for the third element
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			}
		}
	}

	return results
}

func ThreeSumV2(nums []int) [][]int {
	var results [][]int
	length := len(nums)
	sort.Ints(nums)
	for i := 0; i < length-2; i++ {
		num := nums[i]
		//invalid
		if num > 0 {
			break
		}
		//skip duplicate
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := length - 1
		for left < right {
			total := nums[i] + nums[left] + nums[right]
			if total < 0 {
				left += 1
			} else if total > 0 {
				right -= 1
			} else {
				results = append(results, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left += 1
				}
				for left < right && nums[right] == nums[right-1] {
					right -= 1
				}
				left += 1
				right -= 1
			}
		}
	}

	return results
}

func ThreeSumV3(nums []int) [][]int {
	var results [][]int
	sort.Ints(nums)
	length := len(nums)

	for i := 0; i < length-2; i++ {
		num := nums[i]
		//invalid data
		if num > 0 {
			break
		}
		//remove duplicate
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := length - 1
		for left < right {
			total := nums[i] + nums[left] + nums[right]
			if total < 0 {
				left++
			} else if total > 0 {
				right--
			} else {
				//
				results = append(results, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}

	return results
}
