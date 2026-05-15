package main

import (
	"fmt"
	"testing"
)

func Test_findTop3BySum(t *testing.T) {
	stocks := []string{"AAPL", "GOOGL", "MSFT", "AMZN", "TSLA"}

	prices := [][]float32{
		{150.0, 160.0, 140.0, 155.0},     // AAPL: top 3 are 160, 155, 150 (Sum: 465)
		{2800.0, 2900.0, 2700.0, 2850.0}, // GOOGL: top 3 are 2900, 2850, 2800 (Sum: 8550)
		{300.0, 310.0, 305.0, 290.0},     // MSFT: top 3 are 310, 305, 300 (Sum: 915)
		{3400.0, 3300.0, 3500.0, 3450.0}, // AMZN: top 3 are 3500, 3450, 3400 (Sum: 10350)
		{700.0, 800.0, 750.0, 720.0},     // TSLA: top 3 are 800, 750, 720 (Sum: 2270)
	}

	fmt.Println("[Go Variant]#####################")
	findTop3BySum(stocks, prices)
}
