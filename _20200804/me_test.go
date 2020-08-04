package main

import (
	"fmt"
	"testing"
)

func TestCanFinish(t *testing.T) {
	ans1 := CanFinish(2, [][]int{{1, 0}})
	fmt.Println("ans >>> ", ans1)

	ans2 := CanFinish(2, [][]int{{1, 0}, {0, 1}})
	fmt.Println(" ans >>> ", ans2)

	ans3 := CanFinish(7, [][]int{{1, 0}, {0, 3}, {0, 2}, {3, 2}, {2, 5}, {4, 5}, {5, 6}, {2, 4}})
	fmt.Println("ans >>> ", ans3)

	ans4 := CanFinish(3, [][]int{{1, 0}, {1, 2}, {0, 1}})
	fmt.Println("ans >>> ", ans4)

}
