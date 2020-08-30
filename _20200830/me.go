package main

import "strings"

/**
557. 反转字符串中的单词 III

给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。

示例1：
```
输入："Let's take LeetCode contest"
输出："s'teL ekat edoCteeL tsetnoc"
```

提示：
- `在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。`
*/

/**
...
*/
func ReverseWords(s string) string {

	if len(s) == 0 {
		return ""
	}
	arr := strings.Split(s, " ")
	ans := ""
	for i := 0; i < len(arr); i++ {
		ans += reverseString(arr[i]) + " "
	}
	return ans[:len(ans)-1]
}
func reverseString(s string) string {
	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}
