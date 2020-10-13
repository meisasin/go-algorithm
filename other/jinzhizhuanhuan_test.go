package other

import (
	"fmt"
	"testing"
)

func TestJinZhiConversion(t *testing.T) {

	Conversion(1348, 2)
	Conversion(1348, 8)
	Conversion(1348, 16)

}

/**
@param n : 要转换的数值
@param d : 要转换成的进制数
*/
func Conversion(n int, d int) {
	var stack []int
	for n != 0 {
		stack = append(stack, n%d)
		n /= d
	}
	// 输出 START
	for i := len(stack) - 1; i >= 0; i-- {
		fmt.Printf("%d ", stack[i])
	}
	fmt.Println()
	// 输出 END
}
