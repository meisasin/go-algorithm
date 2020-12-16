package main

import "strings"

/**
290. 单词规律

给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。
这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。

示例1：
```
输入: pattern = "abba", str = "dog cat cat dog"
输出: true
```

示例2：
```
输入:pattern = "abba", str = "dog cat cat fish"
输出: false
```

示例3：
```
输入: pattern = "aaaa", str = "dog cat cat dog"
输出: false
```

示例4：
```
输入: pattern = "abba", str = "dog dog dog dog"
输出: false
```

说明:
你可以假设 `pattern` 只包含小写字母， `str` 包含了由单个空格分隔的小写字母。
```
*/

// ------------------------------------------------------------------------------------------

/**

一个月，我终于又回来了
*/

// ------------------------------------------------------------------------------------------
func WordPattern(pattern string, s string) bool {

	charArr := strings.Split(s, " ")

	if len(charArr) != len(pattern) {
		return false
	}

	m := make(map[uint8]string)
	mm := make(map[string]uint8)

	pl := len(pattern)
	for i := 0; i < pl; i++ {
		pv := pattern[i]
		if mv, ok := m[pv]; ok {
			if mv != charArr[i] {
				return false
			}
		} else {
			// 如果 pattern 不存在，那么在 mm 中也不应该存在
			if _, ok := mm[charArr[i]]; ok {
				return false
			}
			m[pv] = charArr[i]
			mm[charArr[i]] = pv
		}

	}

	return true
}
