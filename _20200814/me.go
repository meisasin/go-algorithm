package main

/**
20. 有效的括号

给定一个只包括 `'('，')'，'{'，'}'，'['，']'` 的字符串，判断字符串是否有效。

有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例1：
```
输入: "()"
输出: true
```

示例2：
```
输入: "()[]{}"
输出: true
```

示例3：
```
输入: "(]"
输出: false
```

示例4：
```
输入: "([)]"
输出: false
```

示例5：
```
输入: "{[]}"
输出: true
```
*/

/*
	...
*/
func IsValid(s string) bool {

	sl := len(s)
	if sl%2 != 0 {
		return false
	}
	var stack []byte
	m := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 || m[stack[len(stack)-1]] != s[i] {
				return false
			}
			stack = stack[0 : len(stack)-1]
		}
	}

	if len(stack) > 0 {
		return false
	}
	return true
}
