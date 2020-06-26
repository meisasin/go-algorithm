package main

/**
Pow(x, n)

实现 pow(x, n) ，即计算 x 的 n 次幂函数。

示例 1:
```
输入: 2.00000, 10
输出: 1024.00000
```

示例 2:
```
输入: 2.00000, -2
输出: 0.25000
解释: 2-2 = 1/22 = 1/4 = 0.25
```

说明:

- -100.0 < x < 100.0
- n 是 32 位有符号整数，其数值范围是 [−231, 231 − 1] 。
*/

/**
最近的题目口味怎么变了
*/
func MyPow(x float64, n int) float64 {

	isFu := false
	if n == 0 {
		return 1.0
	}
	if n < 0 {
		isFu = true
		n = n * -1
	}

	res := x

	//tmp := []float64{}
	for i := 1; i < n; i++ {
		res = res * x
	}

	if isFu {
		res = 1 / res
	}
	return res
}
