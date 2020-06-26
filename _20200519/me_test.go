package main

import (
	"fmt"
	"testing"
)

func TestValidPalindrome(t *testing.T) {

	res := ValidPalindrome("abc")
	fmt.Println("Res is ", res)

	res = ValidPalindrome("abca")
	fmt.Println("Res is ", res)
}
