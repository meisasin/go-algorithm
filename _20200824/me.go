package main

/**
459. 重复的子字符串

给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。

示例1：
```
输入: "abab"
输出: True
解释: 可由子字符串 "ab" 重复两次构成。
```

示例2：
```
输入: "aba"
输出: False
```

示例3：
```
输入: "abcabcabcabc"
输出: True
解释: 可由子字符串 "abc" 重复四次构成。 (或者子字符串 "abcabc" 重复两次构成。)

```
*/

/**
这有点不像简单题，看起来简单，做起来还是挺复杂的 >>> 看了题解后，发现自己没救了
*/
func RepeatedSubstringPattern(s string) bool {

	sl := len(s)
	if sl == 0 {
		return true
	}
	first := s[0]
	var idx []int
	for i := 1; i < sl; i++ {
		if s[i] == first {
			idx = append(idx, i)
		}
	}
	if len(idx) == 0 {
		return false
	}
	idx = append(idx, sl)

	for i := 0; i < len(idx); i++ {
		minVal := s[0:idx[i]]
		ml := len(minVal)
		if ml*2 > sl {
			break
		}
		if sl%ml != 0 {
			continue
		}
		ok := true
		begin, end := 0, ml
		for end <= sl {
			if s[begin:end] != minVal {
				ok = false
				break
			}
			begin, end = begin+ml, end+ml
		}
		if ok {
			return true
		}
	}
	return false
}
