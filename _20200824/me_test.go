package main

import (
	"fmt"
	"testing"
)

func TestRepeatedSubstringPattern(t *testing.T) {

	fmt.Println(RepeatedSubstringPattern("abac"))
}

func TestRepeatedSubstringPattern_MiaoA(t *testing.T) {

	fmt.Println(RepeatedSubstringPattern_MiaoA("abab"))
	fmt.Println(RepeatedSubstringPattern_MiaoA("aba"))
	fmt.Println(RepeatedSubstringPattern_MiaoA("abcabcabcabc"))
	fmt.Println(RepeatedSubstringPattern_MiaoA("abac"))
}
