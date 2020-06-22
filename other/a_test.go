package other

import (
	"fmt"
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	val := "aaaaaaaa"
	ans := val[0:1]
	fmt.Println(ans)

	fmt.Println(strings.ReplaceAll("abba", "a", "ccc"))
}
