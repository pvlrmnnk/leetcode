package p0121

type BestTimeToBuyAndSellStockSolution func(prices []int) int

func DumbSolution() BestTimeToBuyAndSellStockSolution {
	return func(prices []int) int {
		profit := 0

		for i := 0; i < len(prices)-1; i++ {
			for j := i + 1; j < len(prices); j++ {
				possibleProfit := prices[j] - prices[i]
				if possibleProfit > profit {
					profit = possibleProfit
				}
			}
		}

		return profit
	}
}

func SmartSolution() BestTimeToBuyAndSellStockSolution {
	return func(prices []int) int {
		left, right := 0, 1
		maxProfit, profit := 0, 0

		for right < len(prices) {
			if prices[left] < prices[right] {
				profit = prices[right] - prices[left]
				if profit > maxProfit {
					maxProfit = profit
				}
			} else {
				left = right
			}
			right++
		}

		return maxProfit
	}
}
