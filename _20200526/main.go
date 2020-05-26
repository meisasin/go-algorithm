package main

import "fmt"

func main() {
	i, j := 5, 11
	k := (i + j) >> 1
	fmt.Println(k)
}

/**
二分查找
*/
func findDuplicate(nums []int) int {
	n := len(nums)
	l, r := 1, n-1
	ans := -1
	for l <= r {
		mid := (l + r) >> 1
		cnt := 0
		for i := 0; i < n; i++ {
			if nums[i] <= mid {
				cnt++
			}
		}
		if cnt <= mid {
			l = mid + 1
		} else {
			r = mid - 1
			ans = mid
		}
	}
	return ans
}

/**
快慢指针
*/
func findDuplicate2(nums []int) int {
	slow, fast := 0, 0
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
