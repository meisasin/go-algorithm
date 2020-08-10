package main

/**
696. 计数二进制子串

给定一个字符串 s，计算具有相同数量0和1的非空(连续)子字符串的数量，并且这些子字符串中的所有0和所有1都是组合在一起的。
重复出现的子串要计算它们出现的次数。

示例 1:
```
输入: "00110011"
输出: 6
解释: 有6个子串具有相同数量的连续1和0：“0011”，“01”，“1100”，“10”，“0011” 和 “01”。

请注意，一些重复出现的子串要计算它们出现的次数。

另外，“00110011”不是有效的子串，因为所有的0（和1）没有组合在一起。
```

示例 2:
```
输入: "10101"
输出: 4
解释: 有4个子串：“10”，“01”，“10”，“01”，它们具有相同数量的连续1和0。
```

注意：
- `s.length 在1到50,000之间。`
- `s 只包含“0”或“1”字符。`
*/

/**
有点乱好吧
*/

func CountBinarySubstrings(s string) int {

	zero, one := 0, 0
	s = "2" + s
	ans := 0
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			if s[i] == '0' {
				zero = 1
				if one > 0 {
					ans++
				}
			} else {
				one = 1
				if zero > 0 {
					ans++
				}
			}
		} else {
			if s[i] == '0' {
				zero++
				if one >= zero {
					ans++
				}
			} else {
				one++
				if zero >= one {
					ans++
				}
			}
		}
	}
	return ans
}

/**
用 map 简化一下代码吧
*/
func CountBinarySubstrings_map(s string) int {

	m := make(map[uint8]int)
	s = "2" + s
	ans := 0
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			m[s[i]] = 1
		} else {
			m[s[i]] = m[s[i]] + 1
		}
		if s[i] == '0' && m['1'] >= m['0'] {
			ans++
		}
		if s[i] == '1' && m['0'] >= m['1'] {
			ans++
		}
	}
	return ans
}
