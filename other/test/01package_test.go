package test

import (
	"fmt"
	"go-algorithm/other"
	"testing"
)

func Test01Package(t *testing.T) {

	ans := other.Package_01(10, []int{5, 4, 3, 5, 3}, []int{400, 300, 350, 500, 200})
	fmt.Println("可获得最大黄金储量值为 ", ans)
}
