package main

/**
336. 回文对

给定一组唯一的单词， 找出所有不同的索引对`(i, j)`，使得列表中的两个单词， `words[i] + words[j]` ，可拼接成回文串。

示例 1:
```
输入: ["abcd","dcba","lls","s","sssll"]
输出: [[0,1],[1,0],[3,2],[2,4]]
解释: 可拼接成的回文串为 ["dcbaabcd","abcddcba","slls","llssssll"]
```

示例 2:
```
输入: ["bat","tab","cat"]
输出: [[0,1],[1,0]]
解释: 可拼接成的回文串为 ["battab","tabbat"]
```
*/

/**
暴力超时
*/
func PalindromePairs(words []string) [][]int {

	var ans [][]int
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if i == j {
				continue
			}
			if isPalind(words[i] + words[j]) {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}

func isPalind(s string) bool {

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
