package main

/**
回文数

判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例1：
```
输入: 121
输出: true
```

示例2：
```
输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
```

示例2：
```
输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。
```

进阶:
你能不将整数转为字符串来解决这个问题吗？
```
*/

/**
还是暴力吧，简单，明了
*/
func DailyTemperatures(T []int) []int {

	le := len(T)
	var ans []int
	for i := 0; i < le; i++ {
		count := 0
		for j := i + 1; j < le; j++ {
			if T[j] > T[i] {
				count = j - i
				break
			}
		}
		ans = append(ans, count)
	}
	return ans
}
