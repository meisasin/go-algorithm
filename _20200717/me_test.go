package main

import (
	"fmt"
	"testing"
)

func TestSearchInsert(t *testing.T) {

	ans := SearchInsertWithApi([]int{1, 3, 5, 6}, 7)
	fmt.Println(ans)
}
