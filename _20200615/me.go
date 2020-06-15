package main

/**
最长公共前缀

编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。爬楼梯

示例1：
```
输入: ["flower","flow","flight"]
输出: "fl"
```

示例2：
```
输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
```
*/

/**
不太对啊，一道简单题我写了这么长的代码吗？
*/

func longestCommonPrefixWithMe(strs []string) string {

	// 优化前
	//if len(strs) == 0 {
	//	return ""
	//}
	//minL := len(strs[0])
	//for i := 1 ; i < len(strs) ; i ++ {
	//	if len(strs[i]) < minL {
	//		minL = len(strs[i])
	//	}
	//}
	//
	//ans := ""
	//for i := 0 ; i < minL ; i ++ {
	//	count := 1
	//	by := strs[0][i]
	//	for j := 1 ; j < len(strs) ; j ++ {
	//		if strs[j][i] == by {
	//			count ++
	//		} else {
	//			break
	//		}
	//	}
	//
	//	if count != len(strs) {
	//		break
	//	} else {
	//		ans = ans + string(by)
	//	}
	//}
	//return ans

	// 优化后   --  我感觉也没用多少内存啊
	ans := ""
	if len(strs) == 0 {
		return ans
	}

	for i := 0; i < len(strs[0]); i++ {
		ch := strs[0][i]
		count := 1
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) {
				break
			}
			if strs[j][i] != ch {
				break
			}
			count++
		}
		if count != len(strs) {
			break
		} else {
			ans = ans + string(ch)
		}
	}
	return ans
}
