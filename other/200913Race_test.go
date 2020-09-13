package other

import (
	"fmt"
	"testing"
)

func Test01(t *testing.T) {
	fmt.Println(" >>> ")
	fmt.Println(numSpecial([][]int{{0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 1, 0, 0, 1}, {0, 0, 0, 0, 1, 0, 0, 0}, {1, 0, 0, 0, 1, 0, 0, 0}, {0, 0, 1, 1, 0, 0, 0, 0}}))
}

func numSpecial(mat [][]int) int {

	rcount, ccount := 0, 0

	rl, cl := len(mat), len(mat[0])

	for i := 0; i < rl; i++ {
		cr := 0
		for j := 0; j < cl; j++ {
			cr ^= mat[i][j]
		}
		if cr == 1 {
			ccount++
		}
	}
	fmt.Println("ccount >>> ", ccount)

	for i := 0; i < cl; i++ {
		rr := 0
		for j := 0; j < rl; j++ {
			rr ^= mat[j][i]
		}
		if rr == 1 {
			rcount++
		}
	}
	fmt.Println("rcount >>> ", rcount)
	return min(rcount, ccount)
}

func unhappyFriends(n int, pf [][]int, pairs [][]int) int {

	pr := map[int][]int{}

	for i := 0; i < len(pairs); i++ {
		x, y := pairs[i][0], pairs[i][1]

		var xqw []int
		for t := 0; t < len(pf[x]); t++ {
			if pf[x][t] == y {
				break
			}
			xqw = append(xqw, pf[x][t])
		}

		pr[x] = xqw

		var yqw []int
		for t := 0; t < len(pf[y]); t++ {
			if pf[y][t] == x {
				break
			}
			yqw = append(yqw, pf[y][t])
		}

		pr[y] = yqw
	}

	fmt.Println(pr)

	ans := 0

	for k, v := range pr {
	a:
		for i := 0; i < len(v); i++ {
			qw := v[i]
			qwArr := pr[qw]
			for j := 0; j < len(qwArr); j++ {
				if qwArr[j] == k {
					ans++
					break a
				}
			}
		}

	}
	return ans
}
