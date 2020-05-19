package main

import (
	"fmt"
	"testing"
)

func TestValidPalindromeWithMe(t *testing.T) {

	res := validPalindromeWithMe("abc")
	fmt.Println("Res is ", res)

	res = validPalindromeWithMe("abca")
	fmt.Println("Res is ", res)
}
