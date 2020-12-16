package ms

import (
	"fmt"
	"math"
	"testing"
)

func TestMinArr(t *testing.T) {

	ans := MinArr([]int{2, 3, 1, 2, 4, 3})
	fmt.Println(ans)

	ans = MinArr([]int{3, 1, 1, 1, 1, 7, 8, 9})
	fmt.Println(ans)

	ans = MinArr([]int{3, 1, 3, 2, 1, 7, 8, 7})
	fmt.Println(ans)
	//
	ans = MinArr([]int{3, 1, 1, 1, 1, 7, 8, 9})
	fmt.Println(ans)
	//
	ans = MinArr([]int{3, 2, 1, 4, 0, 10, 8, 9})
	fmt.Println(ans)
}

func MinArr(arr []int) int {

	ans := math.MaxInt32
	al, l, r := len(arr), 0, 0
	count := 0
	p := 0
	for p < al {
		if p != -1 {
			count += arr[p]
		}
		if count == 7 {
			ans = min(ans, r-l+1)
			r++
			p = r
		} else if count > 7 {
			count -= arr[l]
			l++
			p = -1
		} else {
			r++
			p = r
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
