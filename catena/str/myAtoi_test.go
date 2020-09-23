package str

import (
	"fmt"
	"testing"
)

func TestMyAtoi(t *testing.T) {

	fmt.Println(MyAtoi("42"))
}

func MyAtoi(str string) int {

	first := 0
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			first = i
			break
		}
	}
	str = str[first:]
	if len(str) == 0 {
		return 0
	}

	if str[0] == '-' || (str[0] >= '0' && str[0] <= '9') {
		fu := false
		if str[0] == '-' {
			fu = true
			str = str[1:]
		}
		ans := 0
		point := 0
		for point < len(str) && str[point] >= '0' && str[point] <= '9' {
			ans = ans*10 + int(str[point]-'0')
			ans %= 2147483647
			point++
		}
		if fu {
			ans *= -1
		}
		return ans
	}
	return 0
}
