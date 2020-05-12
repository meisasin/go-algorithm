package main

import (
	"fmt"
	"testing"
)

func TestMyPowWithMe(t *testing.T) {

	fmt.Println("Res ", myPowWithMe(2.00000, 10))
	fmt.Println("Res ", myPowWithMe(2.10000, 3))
	fmt.Println("Res ", myPowWithMe(2.00000, -2))

}
