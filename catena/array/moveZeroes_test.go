package array

import (
	"testing"
)

/**
283. 移动零
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例：
```
输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
```
说明:
必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。
*/
func TestMoveZeroes(t *testing.T) {

	//MoveZeroes([]int { 1, 0 })
	MoveZeroes([]int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0})
}

func MoveZeroes(nums []int) {

	nl, point := len(nums), 0
	for i := 0; i < nl; i++ {
		if nums[i] != 0 {
			nums[point] = nums[i]
			point++
		}
	}
	for i := point; i < nl; i++ {
		nums[i] = 0
	}
}
