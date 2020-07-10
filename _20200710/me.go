package main

/**
309. 最佳买卖股票时机含冷冻期

给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​
设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
- 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
- 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。

示例:
```
输入: [1,2,3,0,2]
输出: 3
解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
```
*/

/**
又没写出来
ERROR
*/
func MaxProfit(prices []int) int {

	l := len(prices)
	if l <= 1 {
		return 0
	}
	if l == 2 {
		return max(0, prices[1]-prices[0])
	}

	sell := make([]int, l)
	buy := make([]int, l)
	mid := make([]int, l)
	sell[0], sell[1] = 0, 0
	buy[0], buy[1] = 0, prices[1]-prices[0]
	mid[0], mid[1] = 0, 0

	for i := 2; i < len(prices); i++ {
		mid[i] = sell[i-1]
		buy[i] = mid[i-1]
		sell[i] = buy[i-1] + (prices[i] - prices[i-1])
	}
	return max(buy[l-1], max(mid[l-1], sell[l-1]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
