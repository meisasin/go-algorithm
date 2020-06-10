package main

/**
回文数

判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例1：
```
输入: 121
输出: true
```

示例2：
```
输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
```

示例2：
```
输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。
```

进阶:
你能不将整数转为字符串来解决这个问题吗？
```
*/

/**
之前做过这题
*/
func isPalindromeWithMe(x int) bool {

	if x < 0 {
		return false
	}

	// 转 int 数组 START
	var arr []int

	for x/10 > 0 {
		arr = append(arr, x%10)
		x = x / 10
	}
	arr = append(arr, x)

	l, r := 0, len(arr)-1
	for l < r {
		if arr[l] != arr[r] {
			return false
		}
		l++
		r--
	}
	return true
	// 转 int 数组 END

	// 转 字符串 START
	// s := strconv.Itoa(x)
	// l, r := 0, len(s) - 1
	// for l < r {
	//     if s[l] != s[r] {
	//         return false
	//     }
	//     l ++
	//     r --
	// }
	// return true
	// 转 字符串 END
}
