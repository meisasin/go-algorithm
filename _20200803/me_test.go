package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestAddStrings(t *testing.T) {
	one := strconv.Itoa(0)
	fmt.Println(one)
	fmt.Println(">>> ", AddStrings("0", "0"))
}
