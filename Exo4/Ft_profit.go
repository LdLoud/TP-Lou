package main

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

func main() {
	Ft_profit([]int{7, 1, 5, 3, 6, 4}) // resultat : 5
	// si on achète au jour 1, nous payons 1,
	// et si nous le vendons au 4eme jour, nous gagnons 6, le bénéfice est 6-1
	Ft_profit([]int{7, 6, 4, 3, 1}) // resultat : 0
}
