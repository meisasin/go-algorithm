package other

import (
	"fmt"
	"sort"
)

func B() {
	arr := []int{1, 2, 3}
	mid := 3
	sort.Slice(arr, func(i, j int) bool {
		x := arr[i] - mid
		y := arr[j] - mid
		if x < 0 {
			x = x * -1
		}
		if y < 0 {
			y = y * -1
		}
		if x > y {
			return true
		} else if x == y {
			return arr[i] > arr[j]
		}
		return false
	})

	fmt.Println(arr)
}
