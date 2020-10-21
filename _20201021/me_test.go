package main

import (
	"fmt"
	"testing"
)

func TestIsLongPressedName(t *testing.T) {
	ans := IsLongPressedName("alex", "aaleex")
	fmt.Println(ans)
}
