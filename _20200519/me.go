package main

/**
验证回文字符串 Ⅱ

给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。

示例 1:
```
输入: "aba"
输出: True
```

示例 2:
```
输入: "abca"
输出: True
解释: 你可以删除c字符。
```
*/

/**
左右往中间对比，如果存在不同的，从不同的索引开始，各舍弃前后两个索引生成两个字段串，再进行对比，如果有一个是回文，即是对的
*/
func ValidPalindrome(s string) bool {

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return isPalindrome(s[i+1:len(s)-i]) || isPalindrome(s[i:len(s)-i-1])
		}
	}
	return true
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
