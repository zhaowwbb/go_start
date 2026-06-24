package main

import "sort"

// CombinationSum finds all unique combinations that sum up to the target.
func CombinationSum(candidates []int, target int) [][]int {
	var result [][]int
	// Sort the candidates to optimize and prune the search space early
	sort.Ints(candidates)
	backtrack(&result, []int{}, candidates, target, 0)
	return result
}

func backtrack(result *[][]int, tempList []int, candidates []int, remain int, start int) {
	if remain == 0 {
		// Go slices are references. We must deep copy the slice before appending.
		copyList := make([]int, len(tempList))
		copy(copyList, tempList)
		*result = append(*result, copyList)
		return
	}

	for i := start; i < len(candidates); i++ {
		// Early pruning: since the candidates are sorted, if the current element
		// exceeds the remainder, all subsequent elements will too.
		if candidates[i] > remain {
			break
		}

		tempList = append(tempList, candidates[i])
		// Pass 'i' instead of 'i + 1' because numbers can be reused
		backtrack(result, tempList, candidates, remain-candidates[i], i)
		// Backtrack: slice off the last element before testing the next branch
		tempList = tempList[:len(tempList)-1]
	}
}

func CombinationSumV2(candidates []int, target int) [][]int {
	var result [][]int
	sort.Ints(candidates)
	backtraceV2(&result, []int{}, candidates, target, 0)
	return result
}

func backtraceV2(result *[][]int, tempList []int, candidates []int, remain int, start int) {
	if remain == 0 {
		copyList := make([]int, len(tempList))
		copy(copyList, tempList)
		*result = append(*result, copyList)
		return
	}

	for i := start; i < len(candidates); i++ {
		if candidates[i] > remain {
			break
		}
		tempList = append(tempList, candidates[i])
		backtraceV2(result, tempList, candidates, remain-candidates[i], i)
		tempList = tempList[:len(tempList)-1]
	}
}

func CombinationSumV3(candidates []int, target int) [][]int {
	var result [][]int
	sort.Ints(candidates)
	backtraceV3(&result, []int{}, candidates, target, 0)
	return result
}

func backtraceV3(result *[][]int, tempList []int, candidates []int, remain int, start int) {
	if remain == 0 {
		copyList := make([]int, len(tempList))
		copy(copyList, tempList)
		*result = append(*result, copyList)
		return
	}

	for i := start; i < len(candidates); i++ {
		if candidates[i] > remain {
			break
		}
		tempList = append(tempList, candidates[i])
		backtraceV3(result, tempList, candidates, remain-candidates[i], i)
		tempList = tempList[:len(tempList)-1]
	}
}
