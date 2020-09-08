package main

/**
77. 组合

给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

示例:
```
输入: n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
```
*/

/**
整懵了
ERROR
*/
func Combine(n int, k int) (ans [][]int) {

	var res [][]int
	for i := 1; i <= n; i++ {

		first, last := i+1, i+k-1
		if last == i {
			res = append(res, []int{i})
			continue
		}
		for last <= n {
			var curr []int
			curr = append(curr, i)

			for j := first; j < last; j++ {
				curr = append(curr, j)
			}
			for j := last; j <= n; j++ {
				var c []int
				c = append(c, curr...)
				c = append(c, j)
				res = append(res, c)
			}
			first++
			last++
		}
	}

	return res
}
