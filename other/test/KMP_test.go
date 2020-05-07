package test

import (
	"fmt"
	"go-algorithm/other"
	"testing"
)

func TestPrefixTable(t *testing.T) {

	prefixTable := other.PrefixTable("ababc")
	fmt.Print(prefixTable)

}

func TestKMPSerach(t *testing.T) {

	idxArr := other.KMPSearch("abaacababcac", "ababc")
	// [5]
	fmt.Println(idxArr)

	fmt.Println("---------------------------------------------------------------")
	idxArr = other.KMPSearch("aaaaab", "aaaa")
	// [0 1]
	fmt.Println(idxArr)

	fmt.Println("---------------------------------------------------------------")
	idxArr = other.KMPSearch("aabaabaaaabbaabaabaaabbaabaabb", "aabb")
	// [8 19 26]
	fmt.Println(idxArr)

	fmt.Println("---------------------------------------------------------------")
	idxArr = other.KMPSearch("abc", "e")
	// []
	fmt.Println(idxArr)
}

func TestGetNext(t *testing.T) {

	prefixTable := other.GetNext("ababc")
	fmt.Print(prefixTable)

}

func TestKMP(t *testing.T) {

	idxArr := other.Kmp("abaacababcac", "ababc", other.GetNext("ababc"))
	fmt.Println(idxArr)

	fmt.Println("---------------------------------------------------------------")
	idxArr = other.Kmp("aaaaab", "aaaa", other.GetNext("aaaa"))
	fmt.Println(idxArr)
}
