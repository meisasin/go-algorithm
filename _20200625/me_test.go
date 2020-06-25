package main

import (
	"fmt"
	"testing"
)

func TestWordBreakWithDFS(t *testing.T) {

	fmt.Println(WordBreakWithDFS("applepenapple", []string{"apple", "pen"})) // true
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~`")
	fmt.Println(WordBreakWithDFS("leetcode", []string{"leet", "code"})) // true
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~`")
	fmt.Println(WordBreakWithDFS("catsandog", []string{"cats", "dog", "sand", "and", "cat"})) // false
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~`")
	fmt.Println(WordBreakWithDFS("goalspecial", []string{"go", "goal", "goals", "special"})) // true
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~`")
	fmt.Println(WordBreakWithDFS("cbca", []string{"bc", "ca"})) // false  c_a
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~`")
	fmt.Println(WordBreakWithDFS("ccaccc", []string{"cc", "ac"})) // true  c_a
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~`")
	fmt.Println(WordBreakWithDFS("a", []string{"a"})) // true

}
