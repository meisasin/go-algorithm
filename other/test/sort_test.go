package test

import (
	"fmt"
	"sort"
	"testing"
)

func TestSimple(t *testing.T) {

	arr := []int{7, 16, 3, 11, 19, 12, 99, 3, 4, 1}

	slice := sort.IntSlice(arr)

	sort.Sort(sort.Reverse(slice))
	fmt.Println(slice)

}

func TestA(t *testing.T) {

	i := rangeBitwiseAnd(1789971928, 1789979928)
	fmt.Println(i)
	fmt.Println("----------------")
	i1 := rangeBitwiseAnd1(1789971928, 1789979928)
	fmt.Println(i1)
}

func rangeBitwiseAnd(m int, n int) int {

	ans := n
	idx := 0
	count := 0
	for ans > m {
		ans &= ^(1 << idx)
		idx++
		count++
	}
	fmt.Println("count:", count)
	return ans
}

func rangeBitwiseAnd1(m int, n int) int {

	ans := n
	count := 0
	for ans > m {
		ans &= ans - 1
		count++
	}
	fmt.Println("count:", count)
	return ans
}
