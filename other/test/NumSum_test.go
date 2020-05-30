package test

import (
	"fmt"
	"go-algorithm/other"
	"testing"
)

func TestTwoNum(t *testing.T) {

	arr := other.TwoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(arr)

}
