package main

/**
763. 划分字母区间

字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一个字母只会出现在其中的一个片段。
返回一个表示每个字符串片段的长度的列表。

示例1：
```
输入：S = "ababcbacadefegdehijhklij"
输出：[9,7,8]
解释：
划分结果为 "ababcbaca", "defegde", "hijhklij"。
每个字母最多出现在一个片段中。
像 "ababcbacadefegde", "hijhklij" 的划分是错误的，因为划分的片段数较少。
```

提示：
- S的长度在[1, 500]之间。
- S只包含小写字母 'a' 到 'z' 。
*/

// ------------------------------------------------------------------------------------------

/**
...
*/

// ------------------------------------------------------------------------------------------
func PartitionLabels(s string) (partition []int) {

	idxArr := make(map[byte][]int)
	for i := 0; i < len(s); i++ {
		if v, ok := idxArr[s[i]]; ok {
			v[1] = i
		} else {
			idxArr[s[i]] = []int{i, i}
		}
	}

	bIdx := idxArr[s[0]]
	begin, end := bIdx[0], bIdx[1]
	var ans []int
	for i := 0; i < len(s); i++ {
		v, _ := idxArr[s[i]]
		if v[1] > end {
			end = v[1]
		} else if i == end {
			ans = append(ans, end-begin+1)
			if i+1 == len(s) {
				break
			}
			bIdx = idxArr[s[i+1]]
			begin, end = bIdx[0], bIdx[1]
		}
	}
	return ans
}
