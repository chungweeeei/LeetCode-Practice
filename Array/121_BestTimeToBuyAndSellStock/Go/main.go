package main

import "fmt"

func maxProfit(prices []int) int {

	if len(prices) < 2 {
		return 0
	}

	maxProfit := 0
	left := 0

	for right := range prices {

		if prices[right] < prices[left] {
			left = right
		} else {
			profit := prices[right] - prices[left]
			maxProfit = max(profit, maxProfit)
		}

	}

	return maxProfit
}

func main() {

	prices := []int{7, 1, 5, 3, 6, 4}

	fmt.Println(maxProfit(prices))

}
