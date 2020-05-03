package main

import (
	"fmt"
	"testing"
)

func TestMaxsubarraywithme(t *testing.T) {
	// [-2,1,-3,4,-1,2,1,-5,4]
	res := maxSubArrayWithMe([]int { -2, 1, -3, 4, -1, 2, 1, -5, 4 })
	fmt.Println(res)
}
