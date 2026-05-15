package main

import (
	"container/heap"
	"fmt"
)

// FloatHeap implements heap.Interface and holds floats.
// This acts as our Min-Heap (standard behavior in Go heap).
type FloatHeap []float32

func (h FloatHeap) Len() int           { return len(h) }
func (h FloatHeap) Less(i, j int) bool { return h[i] < h[j] } // Min-heap logic
func (h FloatHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *FloatHeap) Push(x interface{}) {
	*h = append(*h, x.(float32))
}

func (h *FloatHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// StockResult represents the final structured data for a stock.
// Equivalent to static class StockPriceV4 in Java.
type StockResult struct {
	Name     string
	TotalSum float32
}

// calculateTop3Sum uses a min-heap to keep track of the 3 largest prices.
// Equivalent to calculateTop3SumV4.
func calculateTop3Sum(prices []float32) float32 {
	h := &FloatHeap{}
	heap.Init(h)

	for _, price := range prices {
		heap.Push(h, price)
		// If size exceeds 3, pop the smallest element out of the heap
		if h.Len() > 3 {
			heap.Pop(h)
		}
	}

	// Sum up the remaining 3 elements inside the min-heap
	var sum float32
	for h.Len() > 0 {
		sum += heap.Pop(h).(float32)
	}
	return sum
}

// findTop3BySum matches the behavioral flow of findTop3BySumV4
func findTop3BySum(names []string, prices [][]float32) {
	results := make([]StockResult, 0, len(names))

	for i := 0; i < len(names); i++ {
		sum := calculateTop3Sum(prices[i])
		results = append(results, StockResult{Name: names[i], TotalSum: sum})
	}

	// Quick sort slice in descending order using custom conditional swap inline
	// Equivalent to results.sort((s1, s2) -> Float.compare(s2.totalSum, s1.totalSum))
	for i := 0; i < len(results); i++ {
		for j := i + 1; j < len(results); j++ {
			if results[i].TotalSum < results[j].TotalSum {
				results[i], results[j] = results[j], results[i]
			}
		}
	}

	fmt.Println("Top 3 Stocks by Sum of their Top 3 Prices:")
	limit := 3
	if len(results) < limit {
		limit = len(results)
	}

	for i := 0; i < limit; i++ {
		fmt.Printf("%d. %s - Total Sum: %.2f\n", i+1, results[i].Name, results[i].TotalSum)
	}
}
