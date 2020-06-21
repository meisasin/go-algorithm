package main

import "strings"

/**
验证回文串

给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
说明：本题中，我们将空字符串定义为有效的回文串。

示例1：
```
输入: "A man, a plan, a canal: Panama"
输出: true
```

示例2：
```
输入: "race a car"
输出: false
```
*/

/**
简单题，代码写的确实有点长了，一点都不优雅，一会去找找 Go 自带的校验和转换方法，就不用自己写方法了
*/

func isPalindromeWithMe(s string) bool {

	s = strings.ToLower(s)
	l, r := 0, len(s)-1

	for l < r {
		for l < r && !vaild(s[l]) {
			l++
		}
		for l < r && !vaild(s[r]) {
			r--
		}
		if s[l] != s[r] {
			return false

		}
		l++
		r--
	}
	return true
}

func vaild(u uint8) bool {
	if (u >= 'a' && u <= 'z') || (u >= '0' && u <= '9') {
		return true
	}
	return false
}
