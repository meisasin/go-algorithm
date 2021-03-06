package main

import (
	"fmt"
	"sort"
)

/**
347. 前 K 个高频元素

给定一个非空的整数数组，返回其中出现频率前 k 高的元素。

示例 1:
```
输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
```

示例 2:
```
输入: nums = [1], k = 1
输出: [1]
```

提示：
- `你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。`
- `你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。`
- `题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的。`
- `你可以按任意顺序返回答案。`
*/

/**
这题让我给写的
*/
func TopKFrequent(nums []int, k int) []int {

	m := make(map[int]int)
	nl := len(nums)
	for i := 0; i < nl; i++ {
		m[nums[i]] = m[nums[i]] + 1
	}

	var ans []int
	for _, v := range m {
		ans = append(ans, v)
	}
	sort.Ints(ans)
	ans = ans[len(ans)-k:]
	cm := make(map[int]int)
	for i := 0; i < len(ans); i++ {
		cm[ans[i]] = i
	}

	var realAns []int
	for kk, v := range m {
		if _, ok := cm[v]; ok {
			realAns = append(realAns, kk)
			if len(realAns) == k {
				return realAns
			}
		}
	}
	return realAns
}

func topKFrequenaat(nums []int, k int) []int {

	nl := len(nums)
	m := map[int]int{} // key: 元素值  value: 数量
	for i := 0; i < nl; i++ {
		m[nums[i]]++
	}
	fmt.Println(m)

	res := make([][]int, nl+1)
	rl := len(res)
	for k, v := range m {
		res[v] = append(res[v], k)
	}

	var ans []int
	for i := rl - 1; i >= 0; i-- {
		nn := res[i]
		if len(nn) > 0 {
			for x := 0; x < len(nn); x++ {
				ans = append(ans, nn[x])
				if len(ans) == k {
					return ans
				}
			}
		}
	}
	return ans
}
