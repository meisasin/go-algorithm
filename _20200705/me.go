package main

import "strings"

/**
通配符匹配

给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。
```
'?' 可以匹配任何单个字符。
'*' 可以匹配任意字符串（包括空字符串）。
```

两个字符串完全匹配才算匹配成功。

说明:
- `s` 可能为空，且只包含从 `a-z` 的小写字母。
- `p` 可能为空，且只包含从 `a-z` 的小写字母，以及字符 `?` 和 `*`。

示例1：
```
输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
```

示例2：
```
输入:
s = "aa"
p = "*"
输出: true
解释: '*' 可以匹配任意字符串。
```

示例3：
```
输入:
s = "cb"
p = "?a"
输出: false
解释: '?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。
```

示例4：
```
输入:
s = "adceb"
p = "*a*b"
输出: true
解释: 第一个 '*' 可以匹配空字符串, 第二个 '*' 可以匹配字符串 "dce".
```

示例5：
```
输入:
s = "acdcb"
p = "a*c?b"
输出: false
```

示例6：
 ```
 输入:
 s = "aaaa"
 p = "***a"
 输出: true
```
*/

/**
果然，还是逃脱不了超时的命运
*/
func IsMatch(s string, p string) bool {
	if p == "*" {
		return true
	}
	if p == "?" {
		return len(s) == 1
	}
	if len(s) == 0 && len(p) == 0 {
		return true
	}

	if len(s) == 0 || len(p) == 0 {
		return false
	}
	i, j := 0, 0
	for i < len(s) && j < len(p) {
		if p[j] == '?' {
			i++
			j++
		} else if p[j] == '*' {
			from := j + 1
			for from < len(p) {
				if p[from] != '*' {
					break
				}
				from++
			}
			if from == len(p) {
				return true
			}

			ps := p[from:]
			for x := i; x < len(s); x++ {
				if (ps[0] == '?' || s[x] == ps[0]) && IsMatch(s[x:], ps) {
					return true
				}
			}
			return false
		} else if s[i] == p[j] {
			i++
			j++
		} else {
			return false
		}
	}
	if i == len(s) && j == len(p) {
		return true
	}
	if j == len(p) {
		return false
	}
	return len(strings.ReplaceAll(p[j:], "*", "")) == 0
}
