package main

import (
	"strconv"
)

/**
二进制求和

给你两个二进制字符串，返回它们的和（用二进制表示）。
输入为 非空 字符串且只包含数字 `1` 和 `0`。

示例1：
```
输入: a = "11", b = "1"
输出: "100"
```

示例2：
```
输入: a = "1010", b = "1011"
输出: "10101"
```
提示：
- 每个字符串仅由字符 `'0'` 或 `'1'` 组成。
- `1 <= a.length, b.length <= 10^4`
- 字符串如果不是 `"0"` ，就都不含前导零。
*/

func addBinaryWithMe(a string, b string) string {

	maxL := max(len(a), len(b))

	ans := ""
	jin := false
	for i := 0; i < maxL; i++ {
		count := 0
		if i < len(a) {
			count += int(a[len(a)-1-i] - '0')
		}
		if i < len(b) {
			count += int(b[len(b)-1-i] - '0')
		}
		if jin {
			count += 1
		}
		if count >= 2 {
			jin = true
			count = count - 2
		} else {
			jin = false
		}
		ans = strconv.Itoa(count) + ans
	}
	if jin {
		ans = "1" + ans
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
