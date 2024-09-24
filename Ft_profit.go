package main

import (
	"fmt"
)

func Ft_profit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
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