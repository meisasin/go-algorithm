package main

import (
	"fmt"
	"strings"
)

/**
单词拆分

给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

说明：
- 拆分时可以重复使用字典中的单词。
- 你可以假设字典中没有重复的单词。


示例1：
```
输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
```

示例2：
```
输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
     注意你可以重复使用字典中的单词。
```

示例3：
```
输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false
```
*/

/**
DPS
*/
func WordBreakWithDFS(s string, wordDict []string) bool {

	m := make(map[int]int)
	stack := []int{0}

	wordMap := make(map[string]int)
	for _, word := range wordDict {
		wordMap[word] = 1
	}
	wordMap[""] = 1
	for len(stack) > 0 {
		lastIdx := stack[0]
		stack = stack[1:]
		if _, idxOk := m[lastIdx]; idxOk {
			continue
		}
		m[lastIdx] = 1

		for i := lastIdx + 1; i <= len(s); i++ {
			if _, ok := wordMap[s[lastIdx:i]]; ok {
				if _, otherOk := wordMap[s[i:]]; otherOk {
					return true
				}
				stack = append(stack, i)
			}
		}
	}

	return false
}

/**
ERROR
*/
func WordBreak(s string, wordDict []string) bool {

	var stack []string
	m := make(map[string]int)
	wordMap := make(map[string]int)
	for idx, word := range wordDict {
		wordMap[word] = idx
	}
	begin, end := -1, len(s)
	first := true
	for first || len(stack) > 0 {
		if first {
			begin = 0
		} else {
			lastWord := stack[0]
			idx, _ := m[lastWord]
			begin = idx + 1
			stack = stack[1:]
		}
		first = false
		hasWord := false
		for i := begin + 1; i <= end; i++ {
			if _, ok := wordMap[s[begin:i]]; ok {
				if i == end {
					return true
				}
				if _, ok := wordMap[s[i:]]; ok {
					return true
				}
				word := s[begin:i]
				stack = append(stack, word)
				m[word] = i
				hasWord = true
				//begin = i
			} else {
				if hasWord {
					begin = i
					hasWord = false
				}
			}
		}
		fmt.Println(stack)
	}

	return false
}

/**
ERROR
*/
func WordBreakWithMap(s string, wordDict []string) bool {

	for i := 0; i < len(wordDict); i++ {
		origin := strings.ReplaceAll(s, wordDict[i], "_")
		for j := 0; j < len(wordDict); j++ {
			if i == j {
				continue
			}
			origin = strings.ReplaceAll(origin, wordDict[j], "_")
			if len(strings.ReplaceAll(origin, "_", "")) == 0 {
				return true
			}
		}
	}
	return false
}
