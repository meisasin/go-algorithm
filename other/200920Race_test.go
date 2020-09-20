package other

import (
	"fmt"
	"strings"
	"testing"
)

func TestAA(t *testing.T) {

	//fmt.Println(ReorderSpaces("  this   is  a sentence "))
	fmt.Println(ReorderSpaces("a"))
}

func TestBB(t *testing.T) {

	wordArr := strings.Split("  this   is  a sentence ", " ")
	var realWordArr []string
	wordLen := 0
	for i := 0; i < len(wordArr); i++ {
		if len(wordArr[i]) > 0 && wordArr[i][0] != ' ' {
			realWordArr = append(realWordArr, wordArr[i])
			wordLen += len(wordArr[i])
		}
	}

}

func ReorderSpaces(text string) string {

	tl := len(text)
	wordArr := strings.Split(text, " ")
	var realWordArr []string
	wordLen := 0
	for i := 0; i < len(wordArr); i++ {
		if len(wordArr[i]) > 0 && wordArr[i][0] != ' ' {
			realWordArr = append(realWordArr, wordArr[i])
			wordLen += len(wordArr[i])
		}
	}

	emptyLen := tl - wordLen

	ans := ""
	yu := emptyLen
	al := len(realWordArr)
	if al > 1 {
		num := emptyLen / (al - 1)
		sep := ""
		yu -= num * (al - 1)
		for i := 0; i < num; i++ {
			sep += " "
		}
		for i := 0; i < al; i++ {
			ans = ans + realWordArr[i]
			if i != al-1 {
				ans += sep
			}
		}
	} else {
		ans = realWordArr[0]
	}

	for i := 0; i < yu; i++ {
		ans += " "
	}

	return "-" + ans + "-"
}

func TestCc(t *testing.T) {

	//fmt.Println(maxProductPath([][]int {{4,-3,0,-4,2,0,4},{3,-3,-4,2,-2,-3,-4},{-4,-4,-1,-2,-1,-2,4},{3,-4,-4,2,-4,-4,-2},{1,0,-3,-2,0,2,-1},{2,-4,0,-1,-2,-2,-3}}))
	fmt.Println(maxProductPath([][]int{{-1, 3, 0}, {3, -2, 3}, {-1, 1, -4}}))

}
func maxProductPath(grid [][]int) int {

	x := []int{1, 0}
	y := []int{0, 1}

	rl, cl := len(grid), len(grid[0])
	if rl == 0 || cl == 0 {
		return -1
	}

	gridW := make([][][]int, rl)
	for i := 0; i < rl; i++ {
		gridW[i] = make([][]int, cl)
	}

	var stack [][]int
	stack = append(stack, []int{0, 0})
	gridW[0][0] = append(gridW[0][0], grid[0][0])

	for len(stack) > 0 {
		sl := len(stack)
		fmt.Println(">>>.")
		for i := 0; i < sl; i++ {
			point := stack[i]
			cx, cy := point[0], point[1]

			for j := 0; j < 2; j++ {
				fmt.Println(">>>>>>>.")

				nx, ny := point[0]+x[j], point[1]+y[j]
				nextPoint := []int{nx, ny}
				if nx >= 0 && nx < rl && ny >= 0 && ny < cl {
					nVal := grid[nx][ny]

					for x := 0; x < len(gridW[cx][cy]); x++ {
						fmt.Println("Len >>> ", len(gridW[cx][cy]))
						gridW[nx][ny] = append(gridW[nx][ny], gridW[cx][cy][x]*nVal)
					}
					// 处理一个要不然太多了
					fu := 1
					zheng := -1
					zero := false
					for i := 0; i < len(gridW[nx][ny]); i++ {
						if gridW[nx][ny][i] > 0 {
							zheng = max(gridW[nx][ny][i], zheng)
						} else if gridW[nx][ny][i] < 0 {
							fu = min(gridW[nx][ny][i], fu)
						} else {
							zero = true
						}
					}
					gridW[nx][ny] = []int{}
					if fu != 1 {
						gridW[nx][ny] = append(gridW[nx][ny], fu)
					}
					if zheng != -1 {
						gridW[nx][ny] = append(gridW[nx][ny], zheng)
					}
					if zero {
						gridW[nx][ny] = append(gridW[nx][ny], 0)
					}

					stack = append(stack, nextPoint)
				}
			}
		}
		stack = stack[sl:]
	}

	fmt.Println(gridW)
	ans := -1
	ress := gridW[rl-1][cl-1]
	for i := 0; i < len(ress); i++ {
		ans = max(ans, ress[i])
	}

	if ans < 0 {
		return -1
	}
	return ans % 1000000007
}
