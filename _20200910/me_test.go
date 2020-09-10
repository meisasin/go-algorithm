package main

import (
	"fmt"
	"testing"
)

func TestCombinationSum2(t *testing.T) {

	ans := CombinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	fmt.Println(ans)

	ans1 := CombinationSum2([]int{2, 5, 2, 1, 2}, 5)
	fmt.Println(ans1)

	ans2 := CombinationSum2([]int{14, 6, 25, 9, 30, 20, 33, 34, 28, 30, 16, 12, 31, 9, 9, 12, 34, 16, 25, 32, 8, 7, 30, 12, 33, 20, 21, 29, 24, 17, 27, 34, 11, 17, 30, 6, 32, 21, 27, 17, 16, 8, 24, 12, 12, 28, 11, 33, 10, 32, 22, 13, 34, 18, 12}, 27)
	fmt.Println(ans2)

}
