package other

import (
	"strconv"
	"strings"
	"unicode"
)

/**

字符串解码

给定一个经过编码的字符串，返回它解码后的字符串。
编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

示例:

s = "3[a]2[bc]", 返回 "aaabcbc". aaabcbc
s = "3[a2[c]]", 返回 "accaccacc".
s = "2[abc]3[cd]ef", 返回 "abcabccdcdcdef". abcabccdcdcdef
*/

/**
 */
func DecodeStack(s string) string {
	stack := make([]string, 0)
	num, res := "", ""

	for i := 0; i < len(s); i++ {
		if unicode.IsNumber(rune(s[i])) {
			num += string(s[i])
		} else if s[i] == '[' {
			// 进
			stack = append(stack, num, res)
			num, res = "", ""
		} else if s[i] == ']' {
			context := stack[len(stack)-1:]
			curium, _ := strconv.Atoi((stack[len(stack)-2 : len(stack)-1])[0])
			res = context[0] + strings.Repeat(res, curium)
			// 出
			stack = stack[0 : len(stack)-2]
		} else {
			res += string(s[i])
		}
	}
	return res
}

func Decode(s string) string {

	// 第一个 ] 肯定是和它的 [ 括号离的最近
	for strings.ContainsRune(s, ']') {
		rightIdx := strings.IndexByte(s, ']')
		leftsub := s[0:rightIdx]
		// 从后往前找到最近的一个 '['
		leftIdx := strings.LastIndexByte(leftsub, '[')

		// 数量从 '[' 左边开始往右循环找
		rightNumIdx := leftIdx
		leftNumIdx := strings.LastIndexFunc(leftsub[0:leftIdx-1], func(r rune) bool {
			return !unicode.IsNumber(r)
		}) + 1

		num, _ := strconv.Atoi(leftsub[leftNumIdx:rightNumIdx])

		// 截取是 [) 左开右合，所以从 leftIdx + 1 开始截取
		repeat := strings.Repeat(s[leftIdx+1:rightIdx], num)
		s = s[0:leftNumIdx] + repeat + s[rightIdx+1:]
	}
	return s
}
