package main

import "math"

func main() {

}

var (
	temp []int
	ans  [][]int
)

func findSubsequences(nums []int) [][]int {
	ans = [][]int{}
	dfs(0, math.MinInt32, nums)
	return ans
}

func dfs(cur, last int, nums []int) {
	if cur == len(nums) {
		if len(temp) >= 2 {
			t := make([]int, len(temp))
			copy(t, temp)
			ans = append(ans, t)
		}
		return
	}
	if nums[cur] >= last {
		temp = append(temp, nums[cur])
		dfs(cur+1, nums[cur], nums)
		temp = temp[:len(temp)-1]
	}
	if nums[cur] != last {
		dfs(cur+1, last, nums)
	}
}
