package main

import (
	"fmt"
	"testing"
)

func TestIsMatch(t *testing.T) {
	var ans bool
	//ans := IsMatch("aaaa", "***a")
	//fmt.Println(ans)
	//
	//ans = IsMatch("mississippi", "m??*ss*?i*pi")
	//fmt.Println(ans)

	ans = IsMatch("c", "*?*")
	fmt.Println(ans)

	//ans = IsMatch("a", "a*")
	//fmt.Println(ans)
}
