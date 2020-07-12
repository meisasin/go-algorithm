package other

import (
	"fmt"
	"strings"
)

func reverseStr(s string, k int) string {
	ans := ""

	res := make([]uint8, len(s))
	l, r := 0, min(2*k, len(s))
	for l < len(s) {
		begin, end := l, l+min(k, r-l)
		for i := 0; i < (end+1-begin)/2; i++ {
			res[begin+i], res[end-i] = s[end-i], s[begin+i]
		}
		for i := r - end; i < r; i++ {
			res[i] = s[i]
		}
		l, r = r+1, min(len(s), r+2*k)
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Change() {
	//m := make(map[string]int)

	val := "aaaaaaaa"
	ans := val[0:0]

	fmt.Println(ans)

}

func reformatDate(date string) string {

	monthMap := map[string]string{"Jan": "01", "Feb": "02", "Mar": "03", "Apr": "04", "May": "05", "Jun": "06", "Jul": "07", "Aug": "08", "Sep": "09", "Oct": "10", "Nov": "11", "Dec": "12"}

	res := strings.Split(date, " ")
	ans := res[2] + "-" + monthMap[res[1]] + "-"

	day := ""
	for i := 0; i < len(res[0]); i++ {
		if res[0][i] >= '0' && res[0][i] <= '9' {
			day += string(res[0][i])
		} else {
			break
		}
	}
	if len(day) == 1 {
		day = "0" + day
	}
	fmt.Println(day)
	return ans + day
}
