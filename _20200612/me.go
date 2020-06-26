package main

/**
三数之和

给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
注意：答案中不可以包含重复的三元组。


示例：
```
给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
```
*/

/**
暴力是暴不了力了，通不过啊
*/
import "sort"

func ThreeSum(nums []int) [][]int {

	var ans [][]int

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					cans := []int{nums[i], nums[j], nums[k]}
					sort.Ints(cans)
					equl := false
					for x := 0; x < len(ans); x++ {
						if ans[x][0] == cans[0] && ans[x][1] == cans[1] && ans[x][2] == cans[2] {
							equl = true
							break
						}
					}
					if !equl {
						ans = append(ans, cans)
					}
				}
			}
		}
	}
	return ans
}
