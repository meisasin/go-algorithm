package main

/**
647. 回文子串

给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。
具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。


示例1：
```
输入："abc"
输出：3
解释：三个回文子串: "a", "b", "c"
```

示例2：
```
输入："aaa"
输出：6
解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
```

提示：
- `输入的字符串长度不会超过 1000 。`
*/

/*
	暴力解法 O(n^3)，脑瓜子疼，又来一个算法 >>> Manacher（马拉车）算法:可以在时间复杂度为O(n)的情况下求解一个字符串的最长回文子串长度的问题
*/
func CountSubstrings(s string) int {

	ans := 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isSub(s[i : j+1]) {
				ans++
			}
		}
	}
	return ans
}

/*
	这种方式怎么这么快，有点懵好吧，也是 O(n^2) 的复杂度啊
*/
func CountSubstrings_(s string) int {

	ans := 0
	sl := len(s)
	for i := 0; i < sl*2-1; i++ {
		l, r := i/2, i/2+i%2
		for l >= 0 && r < sl {
			if s[l] != s[r] {
				break
			}
			l--
			r++
			ans++
		}
	}
	return ans
}
func isSub(s string) bool {
	if len(s) == 1 {
		return true
	}
	begin, end := 0, len(s)-1
	for begin < end {
		if s[begin] != s[end] {
			return false
		}
		begin++
		end--
	}
	return true
}
