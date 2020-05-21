package main

/**
最长回文子串

给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1:
```
输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
```

示例 2:
```
输入: "cbbd"
输出: "bb"
```
*/

/**
写的有点长哟~~~ ,好的呢好的呢
*/
func longestPalindromeWithMe(s string) string {
	ml := 0
	mr := -1
	for i := 0; i < len(s); i++ {
		l := i
		r := i
		for l-1 >= 0 && r+1 < len(s) && s[l-1] == s[r+1] {
			l--
			r++
		}
		if r-l > mr-ml {
			ml = l
			mr = r
		}

		l = i + 1
		r = i
		for l-1 >= 0 && r+1 < len(s) && s[l-1] == s[r+1] {
			l--
			r++
		}
		if r-l > mr-ml {
			ml = l
			mr = r
		}
	}
	var max []byte
	for i := ml; i <= mr; i++ {
		max = append(max, s[i])
	}
	return string(max)
}
