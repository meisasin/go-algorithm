package math

import (
	"fmt"
	"math"
	"testing"
)

func TestCountPrimes(t *testing.T) {
	count := CountPrimes(10)
	fmt.Println("COUNT >>> ", count)
}
func CountPrimes(n int) int {
	count := 0
a:
	for i := 2; i <= n; i++ {
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				continue a
			}
		}
		count++
	}
	return count
}
