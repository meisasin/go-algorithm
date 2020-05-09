package main

import "math"

/**
 x 的平方根

实现 `int sqrt(int x)` 函数。

计算并返回 x 的平方根，其中 x 是非负整数。

由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

示例 1:
```
输入: 4
输出: 2
```

示例 2:
```
输入: 8
输出: 2
说明: 8 的平方根是 2.82842...,
     由于返回类型是整数，小数部分将被舍去。
```
*/

/**
就今天这题，没点数学功底的还真做不出来，不由的流下了基础太差的眼泪

	二分查找法本来还想到来着，有点笨了，居然没试试，但是有个很明显的缺陷，没办法精确小数位
*/
func mySqrtWithMe(x int) int {
	ans := int(math.Exp(0.5 * math.Log(float64(x))))
	if (ans+1)*(ans+1) <= x {
		return ans + 1
	}
	return ans
}

func BinarySqrt(x int) int {

	l, r, ans := 0, x, -1

	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}

	}
	return ans
}
