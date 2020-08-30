package other

import (
	"fmt"
	"testing"
)

func TestC(t *testing.T) {
	fmt.Println(containsPattern([]int{2, 2}, 1, 2))
}

func containsPattern(arr []int, m int, k int) bool {

	al := len(arr)
	if m*k > al {
		return false
	}

	stop := al - (m * k)
	for i := 0; i <= stop; i++ {
		count := 1
		l, r := i, i+m
		nl, nr := r, r+m
		for nr <= al {
			ok := true
			for x := 0; x < m; x++ {
				if arr[x+nl] != arr[x+l] {
					ok = false
					break
				}
			}
			if ok {
				count++
			} else {
				break
			}
			nl, nr = nr, nr+m
		}
		if count >= k {
			return true
		}
	}
	return false
}

func TestGetMaxLen(t *testing.T) {
	//fmt.Println(getMaxLen([]int { 1,-2,-3,4 }))
	fmt.Println(getMaxLen([]int{0, 1, -2, -3, -4}))
	//fmt.Println(getMaxLen([]int { -1,-2,-3,0,1 }))
	//fmt.Println(">>> ", getMaxLen([]int { 1,2,3,5,-6,4,0,10 }))
}
func getMaxLen(nums []int) int {

	nums = append(nums, 0)
	nl := len(nums)

	dp := make([]int, nl)
	jdp := make([]int, nl)
	cdp := make([]int, nl)

	maxVal := 0
	for i := nl - 2; i >= 0; i-- {
		if nums[i] == 0 {
			dp[i] = 0
			jdp[i] = 0
			cdp[i] = 0
		} else if nums[i] > 0 {
			if jdp[i+1]%2 == 0 {
				dp[i] = cdp[i+1] + 1
			} else if nums[i+1] < 0 {
				dp[i] = 1
			} else {
				dp[i] = dp[i+1] + 1
			}
			jdp[i] = jdp[i+1]
			cdp[i] = cdp[i+1] + 1
		} else {
			if jdp[i+1]%2 == 1 {
				dp[i] = cdp[i+1] + 1
			} else {
				dp[i] = dp[i+1]
			}
			jdp[i] = jdp[i+1] + 1
			cdp[i] = cdp[i+1] + 1
		}
		maxVal = max(maxVal, dp[i])
		fmt.Println(maxVal)
	}
	return maxVal
}

func reverseString(s string) string {
	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}
