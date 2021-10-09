package main

import "math"

//剑指 Offer 63. 股票的最大利润

//假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？

func maxProfit(prices []int) int {
	min := math.MaxInt64
	res := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i]-min > res {
			res = prices[i] - min
		}

	}
	return res
}
