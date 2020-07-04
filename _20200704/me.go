package main

/**
最长有效括号

给定一个只包含 `'('` 和 `')'` 的字符串，找出最长的包含有效括号的子串的长度。

示例1：
```
输入: "(()"
输出: 2
解释: 最长有效括号子串为 "()"
```

示例2：
```
输入: ")()())"
输出: 4
解释: 最长有效括号子串为 "()()"
```
*/

/**
栈 + 辅助数组
*/
func LongestValidParentheses(s string) int {

	var stack []int
	res := make([]bool, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if len(stack) > 0 {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[idx], res[i] = true, true
			res[i] = true
		}
	}
	max, count := 0, 0
	for i := 0; i < len(res); i++ {
		if res[i] {
			count++
		} else {
			max = maxN(max, count)
			count = 0
		}
	}
	max = maxN(max, count)
	return max
}

func maxN(a, b int) int {
	if a > b {
		return a
	}
	return b
}
