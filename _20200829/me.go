package main

/**
214. 最短回文串

给定一个字符串 s，你可以通过在字符串前面添加字符将其转换为回文串。找到并返回可以用这种方式转换的最短回文串。

示例1：
```
输入: "aacecaaa"
输出: "aaacecaaa"
```

示例2：
```
输入: "abcd"
输出: "dcbabcd"
```
*/

/**
复杂度比较高，俺知足了，毕竟困难题（说起来按这种复杂度做的话，也不算困难题了吧）
*/
func ShortestPalindrome(s string) string {
	pre, suf := "", ""
	for len(s) > 0 {
		if isPalind(s) {
			break
		}
		last := string(s[len(s)-1])
		pre += last
		suf = last + suf
		s = s[:len(s)-1]
	}
	return pre + s + suf
}

func isPalind(s string) bool {
	if len(s) <= 1 {
		return true
	}
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
