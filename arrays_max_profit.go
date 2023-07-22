package leetcode_75

/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
*/

/*
Example:

Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
*/

// Naive solution is O(n^2)
// We want a data structure that is, for i, the MAX value of the array after i.
// [7, 1, 5, 3, 6, 4]
// [6, 6, 6, 6, 4, -1]
// Then subtract.
// So we can go backwards through the array to construct this.
// First element of array -1, yeah, this works.

func maxProfit(prices []int) int {
	maxAfter := make([]int, len(prices))
	var currentMax int = -1
	maxAfter[len(prices) - 1] = currentMax
	for i := len(prices) - 2; i >= 0; i-- {
		if prices[i + 1] > currentMax {
			maxAfter[i] = prices[i + 1]
			currentMax = prices[i + 1]
		} else {
			maxAfter[i] = currentMax
		}
	}
	var maxProfit int
	for i, _ := range maxAfter {
		j := maxAfter[i] - prices[i]
		if j > maxProfit {
			maxProfit = j
		}
	}
	return maxProfit
}
