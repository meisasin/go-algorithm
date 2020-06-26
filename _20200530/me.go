package main

/**
柱状图中最大的矩形

给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

求在该柱状图中，能够勾勒出来的矩形的最大面积。
*/

/**
写懵逼了, 还是得先用暴力解法写，要不然会绕进去，出不来，把暴力解法都忘了
*/
func LargestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	maxArea := heights[0]
	horiz := 1
	min := heights[0]
	minIdx := 0

	for i := 1; i < len(heights); i++ {
		currMin := minNum(min, heights[i])
		oneArea := heights[i]
		twoArea := currMin * (horiz + 1)
		threeArea := maxArea

		if oneArea >= twoArea && oneArea >= threeArea {
			maxArea = oneArea
			horiz = 1
			min = maxArea
			minIdx = i
		} else if twoArea >= threeArea {
			horiz++
			maxArea = twoArea
			min = currMin
			if heights[i] <= min {
				minIdx = i
			}
		} else {
			horiz = 0
			if i+1 < len(heights) {
				min = heights[i+1]
				minIdx = i + 1
			}
		}

		fourArea := heights[i]
		choriz := 1
		cmin := heights[i]
		cIdx := i
		for j := i - 1; j > minIdx; j-- {

			if minNum(cmin, heights[j])*(choriz+1) >= fourArea {
				cmin = minNum(cmin, heights[j])
				choriz++
				cIdx = j
				fourArea = cmin * choriz
			}

		}

		if fourArea >= maxArea {
			maxArea = fourArea
			horiz = choriz
			min = cmin
			minIdx = cIdx
		}
	}
	return maxArea
}

func largestRectangleAreaBaoLi(heights []int) int {

	if len(heights) == 0 {
		return 0
	}
	maxArea := heights[0]

	for i := 0; i < len(heights); i++ {

		l := i - 1
		r := i + 1
		for l >= 0 && heights[l] >= heights[i] {
			l--
		}
		for r < len(heights) && heights[r] >= heights[i] {
			r++
		}

		currArea := (r - l - 1) * heights[i]
		if currArea > maxArea {
			maxArea = currArea
		}
	}
	return maxArea

}

func minNum(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
