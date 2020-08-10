package main

func main() {

}

func countBinarySubstrings(s string) int {
	var ptr, last, ans int
	n := len(s)
	for ptr < n {
		c := s[ptr]
		count := 0
		for ptr < n && s[ptr] == c {
			ptr++
			count++
		}
		ans += min(count, last)
		last = count
	}

	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
