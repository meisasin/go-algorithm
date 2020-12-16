package main

import (
	"fmt"
	"testing"
)

func TestWordPattern(t *testing.T) {

	fmt.Println(WordPattern("abba", "dog cat cat dog"))
	fmt.Println(WordPattern("abba", "dog cat cat fish"))
	fmt.Println(WordPattern("aaaa", "dog cat cat dog"))
	fmt.Println(WordPattern("abba", "dog dog dog dog"))

}
