package icon_pln_logical

import (
	"math"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	maxProfit := func(prices []int) int {
		if len(prices) < 2 {
			return 0
		}

		minPrice := math.MaxInt
		maxProfit := 0

		for _, price := range prices {
			if price < minPrice {
				minPrice = price
			}
			profit := price - minPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		}

		return maxProfit
	}

	tests := []struct {
		prices   []int
		expected int
	}{
		{[]int{10, 9, 6, 5, 15}, 10},
		{[]int{7, 8, 3, 10, 8}, 7},
		{[]int{5, 12, 11, 12, 10}, 7},
		{[]int{7, 18, 27, 10, 29}, 22},
		{[]int{20, 17, 15, 14, 10}, 0},
	}

	for _, tt := range tests {
		got := maxProfit(tt.prices)
		if got != tt.expected {
			t.Errorf("maxProfit(%v) = %d; want %d", tt.prices, got, tt.expected)
		}
	}
}
