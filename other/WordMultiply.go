package other

import "sort"

/**
最大单词长度乘积
给定一个字符串数组 words，找到 length(word[i]) * length(word[j]) 的最大值，
并且这两个单词不含有公共字母。你可以认为每个单词只包含小写字母。如果不存在这样的两个单词，返回 0。

示例 1:

输入: ["abcw","baz","foo","bar","xtfn","abcdef"]
输出: 16
解释: 这两个单词为 "abcw", "xtfn"。
示例 2:

输入: ["a","ab","abc","d","cd","bcd","abcd"]
输出: 4
解释: 这两个单词为 "ab", "cd"。
示例 3:

输入: ["a","aa","aaa","aaaa"]
输出: 0
解释: 不存在这样的两个单词。

*/

/**
计算之前先排序的做法
*/
// ["abcw","baz","foo","bar","xtfn","abcdef"]
func maxProductWithSort(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})

	marks := make([]int, len(words))
	for i, n := range words {
		bitmark1 := 0
		for _, i := range n {
			bitmark1 |= 1 << (i - 'a')
		}
		marks[i] = bitmark1
	}

	max := 0
a:
	for i := 0; i < len(words)-1; i++ {
		if len(words[i])*len(words[i+1]) <= max {
			break
		}
		for j := i + 1; j < len(words); j++ {
			if marks[i]&marks[j] == 0 {
				//if !strings.ContainsAny(words[i], words[j]) {
				if len(words[i])*len(words[j]) > max {
					max = len(words[i]) * len(words[j])
					continue a
				}
			}
		}
	}
	return max
}

/**
计算之前不排序的做法
*/
func MaxProduct(words []string) int {

	marks := make([]int, len(words))
	for i, n := range words {
		for _, b := range n {
			marks[i] |= 1 << (b - 'a')
		}
	}

	max := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if marks[i]&marks[j] == 0 {
				if len(words[i])*len(words[j]) > max {
					max = len(words[i]) * len(words[j])
				}
			}
		}
	}
	return max
}
