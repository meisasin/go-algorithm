package array

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	//source := [][]int { { 1,2,3 } , { 4,5,6 } , { 7,8,9 }}
	source := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	Rotate(source)
	fmt.Println(source)

}

// 我太难了
func Rotate(matrix [][]int) {

	// mid := (len(matrix) - 1) / 2 + 1
	ml := len(matrix)
	for i := 0; i < ml-1; i++ {
		for j := i; j < ml-1-i; j++ {
			currPoint := []int{i, j}
			nextPoint := []int{currPoint[1], ml - 1 - currPoint[0]}
			temp := matrix[currPoint[0]][currPoint[1]]
			temp, matrix[nextPoint[0]][nextPoint[1]] = matrix[nextPoint[0]][nextPoint[1]], temp

			currPoint = nextPoint
			nextPoint = []int{currPoint[1], ml - 1 - currPoint[0]}
			temp, matrix[nextPoint[0]][nextPoint[1]] = matrix[nextPoint[0]][nextPoint[1]], temp

			currPoint = nextPoint
			nextPoint = []int{currPoint[1], ml - 1 - currPoint[0]}
			temp, matrix[nextPoint[0]][nextPoint[1]] = matrix[nextPoint[0]][nextPoint[1]], temp

			currPoint = nextPoint
			nextPoint = []int{currPoint[1], ml - 1 - currPoint[0]}
			temp, matrix[nextPoint[0]][nextPoint[1]] = matrix[nextPoint[0]][nextPoint[1]], temp
		}
	}

}
