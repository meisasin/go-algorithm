package main

import (
	"sort"
	"testing"
)

func TestRespace(t *testing.T) {

	dict := []string{"abc", "deff", "cccccc"}
	sort.Slice(dict, func(i, j int) bool {
		return len(dict[i]) > len(dict[j])
	})

}
