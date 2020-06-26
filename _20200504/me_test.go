package main

import (
	"fmt"
	"testing"
)

func TestJump(t *testing.T) {
	// [-2,1,-3,4,-1,2,1,-5,4]
	res := Jump([]int{2, 3, 1, 1, 4})
	fmt.Println(res)
}
