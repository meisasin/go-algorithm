package main

import (
	"fmt"
	"testing"
)

func TestSumNumsWithMe(t *testing.T) {

	s := "1jkl;"
	smap := map[uint8]bool{s[0]: true}
	//fmt.Println(val)
	if val, ok := smap[s[0]]; ok {
		fmt.Println("Has val ...")
		fmt.Println(val)
	}
	fmt.Println(smap)

	//fmt.Println(val == "")

	//aaa := map[uint8]bool { s[I]: true }
	fmt.Println(len(" "))
	m := " "
	for i := 0; i < len(m); i++ {
		fmt.Println("~~~~")
	}

}
