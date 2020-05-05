package test

import (
	"fmt"
	"go-algorithm/other"
	"testing"
)

/**
判断两个字符串是否重复
*/

func TestCheckNoRepeat(t *testing.T) {

	repeat := other.CheckNoRepeat("aa", "bb")
	fmt.Println(repeat)

	repeat = other.CheckNoRepeat("aab", "bb")
	fmt.Println(repeat)

	repeat = other.CheckNoRepeat("aade", "bb")
	fmt.Println(repeat)

	repeat = other.CheckNoRepeat("debbee", "bba")
	fmt.Println(repeat)

}
