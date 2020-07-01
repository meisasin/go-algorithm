package main

/**
最长重复子数组

给两个整数数组 `A` 和 `B` ，返回两个数组中公共的、长度最长的子数组的长度。

示例1：
```
输入:
A: [1,2,3,2,1]
B: [3,2,1,4,7]
输出: 3
解释:
长度最长的公共子数组是 [3, 2, 1]。
```
*/

/**
先来一发暴力解法
*/
func FindLength(A []int, B []int) int {

	var max int
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			if A[i] == B[j] {
				size := 1
				m, n := i+1, j+1
				for n < len(B) && m < len(A) && A[m] == B[n] {
					size++
					m++
					n++
				}
				max = maxNum(max, size)
			}
		}
	}
	return max
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
DP
*/
func FindLengthWithDP(A []int, B []int) int {

	m, n := len(A), len(B)
	arr := make([][]int, m+1)
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]int, n+1)
	}
	ans := 0
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if A[i] == B[j] {
				arr[i][j] = arr[i+1][j+1] + 1
			} else {
				arr[i][j] = 0
			}
			if ans < arr[i][j] {
				ans = arr[i][j]
			}
		}
	}
	return ans
}
