package test

import (
	"fmt"
	"go-algorithm/other"
	"testing"
	"time"
)

func TestDecode(t *testing.T) {

	// lleft: idx: 3, sub: 3[a[a; right: idx: 5, sub: -
	res := other.Decode("3[a]2[bc]")
	fmt.Printf(" ----> %s\n", res)

	res = other.Decode("2[abc]3[cd]ef")
	fmt.Printf(" ----> %s\n", res)

}

// BenchmarkDecode-12    	1000000000	         0.000009 ns/op
func BenchmarkDecode(b *testing.B) {

	res := other.Decode("2[abc]3[cd]ef")
	fmt.Printf(" ----> %s\n", res)

}

// BenchmarkDecodeStack-12    	1000000000	         0.000011 ns/op
func BenchmarkDecodeStack(b *testing.B) {

	res := other.DecodeStack("2[abc]3[cd]ef")
	fmt.Printf(" ----> %s\n", res)
}

func BenchmarkA(b *testing.B) {

	time.Sleep(time.Nanosecond * 1)

}
