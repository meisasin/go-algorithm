package main

import (
	"fmt"
	"testing"
)

func TestMyPow(t *testing.T) {

	fmt.Println("Res ", MyPow(2.00000, 10))
	fmt.Println("Res ", MyPow(2.10000, 3))
	fmt.Println("Res ", MyPow(2.00000, -2))

}
