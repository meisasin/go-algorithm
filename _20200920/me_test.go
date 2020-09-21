package main

import (
	"fmt"
	"testing"
)

func TestSubsets(t *testing.T) {

}

func TestSubsets_bit(t *testing.T) {

	fmt.Println(subsets_bit([]int{1, 2, 3}))
}

func TestMask(t *testing.T) {

	//i := 7
	//for mask := 0 ; mask < 7 ; mask ++ {
	//	if mask >> i & 1 > 0 {
	//
	//	}
	//}

	//110
	mask := 6
	mover := 0
	fmt.Println(mask >> mover)
	fmt.Println(mask >> mover & 1)

	fmt.Println(7 & 1)
}
