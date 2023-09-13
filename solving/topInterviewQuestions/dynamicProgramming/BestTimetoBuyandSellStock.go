package dynamicProgramming

func BestTimetoBuyandSellStock(prices []int) int {
	var maxPrice, profit, newProfit int
	minPrice := -1
	for i := 0; i < len(prices); i++ {
		if minPrice == -1 || minPrice > prices[i] {
			minPrice = prices[i]
			maxPrice = 0
			newProfit = maxPrice - minPrice
			if newProfit > profit {
				profit = newProfit
			}

			continue
		}

		if prices[i] > maxPrice {
			maxPrice = prices[i]
			newProfit = maxPrice - minPrice
			if newProfit > profit {
				profit = newProfit
			}
		}
	}

	if profit < 0 {
		profit = 0
	}

	return profit
}
