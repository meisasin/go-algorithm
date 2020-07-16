package main

import (
	"fmt"
	"testing"
)

func TestIsBipartite(t *testing.T) {

	ans := IsBipartite([][]int{{1}, {0}, {4}, {4}, {2, 3}})
	fmt.Println(ans)

}
