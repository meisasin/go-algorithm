package main

/**
343. 整数拆分

给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。

示例1：
```
输入: 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1。
```

示例2：
```
输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
```
*/

/**
思路就是 3 * 3 > 2 * 2 * 2  >>> 即每次将数减去 3 直到不够3，最小到4，因为 4 可以分解成 1 * 3 | 2 * 2, 应取 2 * 2
*/
func IntegerBreak(n int) int {

	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	return caseBreak(n)
}

func caseBreak(n int) int {
	if n < 5 {
		return n
	}
	return 3 * caseBreak(n-3)
}
