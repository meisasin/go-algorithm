package other

import "fmt"

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
