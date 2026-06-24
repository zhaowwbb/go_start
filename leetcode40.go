package main

import "sort"

// CombinationSum2 finds all unique combinations where candidate numbers sum to target.
func CombinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	// Sorting is mandatory to group duplicate candidates together
	sort.Ints(candidates)
	backtrack(&result, []int{}, candidates, target, 0)
	return result
}

func backtrack(result *[][]int, tempList []int, candidates []int, remain int, start int) {
	if remain == 0 {
		copyList := make([]int, len(tempList))
		copy(copyList, tempList)
		*result = append(*result, copyList)
		return
	}

	for i := start; i < len(candidates); i++ {
		// Early pruning
		if candidates[i] > remain {
			break
		}

		// Skip duplicate combinations by skipping duplicate adjacent items
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}

		tempList = append(tempList, candidates[i])
		// Pass 'i + 1' so each candidate index is only chosen once per branch path
		backtrack(result, tempList, candidates, remain-candidates[i], i+1)
		tempList = tempList[:len(tempList)-1]
	}
}

func CombinationSum2V2(candidates []int, target int) [][]int {
	var result [][]int
	sort.Ints(candidates)
	backTrace2V2(&result, []int{}, candidates, target, 0)
	return result
}

func backTrace2V2(result *[][]int, tempList []int, candidates []int, remain int, start int) {
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

		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		tempList = append(tempList, candidates[i])
		backTrace2V2(result, tempList, candidates, remain-candidates[i], i+1)
		tempList = tempList[:len(tempList)-1]

	}
}
