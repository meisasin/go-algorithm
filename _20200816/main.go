package main

func main() {

}

var (
	dx = []int{1, 0, 0, -1}
	dy = []int{0, 1, -1, 0}
)

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	currColor := image[sr][sc]
	if currColor != newColor {
		dfs(image, sr, sc, currColor, newColor)
	}
	return image
}

func dfs(image [][]int, x, y, color, newColor int) {
	if image[x][y] == color {
		image[x][y] = newColor
		for i := 0; i < 4; i++ {
			mx, my := x+dx[i], y+dy[i]
			if mx >= 0 && mx < len(image) && my >= 0 && my < len(image[0]) {
				dfs(image, mx, my, color, newColor)
			}
		}
	}
}
