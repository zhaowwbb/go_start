package main

// swap handles the in-place index manipulation using Go's multiple assignment syntax.
func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func swapV4(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

// --- Implementation 1 ---

func partition(nums []int, low, high int) int {
	pivot := nums[high]
	i := low - 1
	for j := low; j < high; j++ {
		if nums[j] < pivot {
			i++
			swap(nums, i, j)
		}
	}
	swap(nums, i+1, high)
	return i + 1
}

func partitionV4(nums []int, low, high int) int {
	pivot := nums[high]
	i := low - 1
	for j := low; j < high; j++ {
		if nums[j] < pivot {
			i++
			swapV4(nums, i, j)
		}
	}
	swapV4(nums, i+1, high)
	return i + 1
}

func quickSortV4(nums []int, low, high int) {
	if low < high {
		pivotIndex := partitionV4(nums, low, high)
		quickSortV4(nums, low, pivotIndex-1)
		quickSortV4(nums, pivotIndex+1, high)
	}
}

func Sort4(nums []int) {
	if len(nums) == 0 {
		return
	}
	quickSortV4(nums, 0, len(nums)-1)
}

func quickSort(nums []int, low, high int) {
	if low < high {
		pivotIndex := partition(nums, low, high)
		quickSort(nums, low, pivotIndex-1)
		quickSort(nums, pivotIndex+1, high)
	}
}

// Sort public entry point for version 1
func Sort(nums []int) {
	if len(nums) == 0 {
		return
	}
	quickSort(nums, 0, len(nums)-1)
}

// --- Implementation 2 ---

func partition2(nums []int, low, high int) int {
	highValue := nums[high]
	i := low - 1
	for j := low; j < high; j++ {
		if nums[j] < highValue {
			i++
			swap(nums, i, j)
		}
	}
	swap(nums, high, i+1)
	return i + 1
}

func quickSort2(nums []int, low, high int) {
	if low < high {
		pivot := partition2(nums, low, high)
		quickSort2(nums, low, pivot-1)
		quickSort2(nums, pivot+1, high)
	}
}

// Sort2 public entry point for version 2
func Sort2(nums []int) {
	quickSort2(nums, 0, len(nums)-1)
}

// --- Implementation 3 ---

func partition3(nums []int, low, high int) int {
	highValue := nums[high]
	i := low - 1
	for j := low; j < high; j++ {
		if nums[j] < highValue {
			i++
			swap(nums, i, j)
		}
	}
	swap(nums, high, i+1)
	return i + 1
}

func quickSort3(nums []int, low, high int) {
	if low < high {
		pivot := partition3(nums, low, high)
		quickSort3(nums, low, pivot-1)
		quickSort3(nums, pivot+1, high)
	}
}

// Sort3 public entry point for version 3
func Sort3(nums []int) {
	quickSort3(nums, 0, len(nums)-1)
}

func swap5(nums []int, low, high int) {
	nums[low], nums[high] = nums[high], nums[low]
}

func partition5(nums []int, low, high int) int {
	pivot := nums[high]
	i := low - 1
	for j := low; j < high; j++ {
		if nums[j] < pivot {
			i++
			swap5(nums, i, j)
		}
	}
	swap5(nums, i+1, high)
	return i + 1
}

func quickSort5(nums []int, low, high int) {
	if low < high {
		pivotIndex := partition5(nums, low, high)
		quickSort5(nums, low, pivotIndex-1)
		quickSort5(nums, pivotIndex+1, high)
	}
}

func Sort5(nums []int) {
	if len(nums) == 0 {
		return
	}
	quickSort5(nums, 0, len(nums)-1)
}
