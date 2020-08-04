package main

/**
415. 字符串相加

你这个学期必须选修 numCourse 门课程，记为 `0` 到 `numCourse-1` 。
在选修某些课程之前需要一些先修课程。 例如，想要学习课程 `0` ，你需要先完成课程 `1` ，我们用一个匹配来表示他们：`[0,1]`
给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？

示例 1:
```
输入: 2, [[1,0]]
输出: true
解释: 总共有 2 门课程。学习课程 1 之前，你需要完成课程 0。所以这是可能的。
```

示例 2:
```
输入: 2, [[1,0],[0,1]]
输出: false
解释: 总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0；并且学习课程 0 之前，你还应先完成课程 1。这是不可能的。
```

提示：
- `输入的先决条件是由 边缘列表 表示的图形，而不是 邻接矩阵 。详情请参见图的表示法。`
- `你可以假定输入的先决条件中没有重复的边。`
- `1 <= numCourses <= 10^5`
*/

/**
果然还是被我的毅力打倒了吧。 话说写完就不认识了，也挺尴尬的 >>> 下次遇见还得想半天 o_o
*/
func CanFinish(nc int, pq [][]int) bool {

	pMap := make(map[int][]int)
	for i := 0; i < len(pq); i++ {
		key, value := pq[i][0], pq[i][1]
		pMap[key] = append(pMap[key], value)
	}

	studyed := make(map[int]int)

	for i := 0; i < nc; i++ {
		waitStudy := make(map[int]int)

		ans := search(pMap, studyed, waitStudy, i)
		if !ans {
			return false
		}
	}
	return true
}

func search(pMap map[int][]int, studyed map[int]int, waitStudy map[int]int, node int) bool {
	// 已经学了，返回 true
	if _, ok := studyed[node]; ok {
		return true
	}
	// 未学，但没方向的，返回 true
	if _, ok := pMap[node]; !ok {
		studyed[node] = 1
		return true
	}
	// 未学，但在待学习中，返回 false
	if _, ok := waitStudy[node]; ok {
		return false
	}
	values := pMap[node]
	okc := 0
	errc := 0
	waitStudy[node] = 1
	for i := 0; i < len(values); i++ {
		nodeAsn := search(pMap, studyed, waitStudy, values[i])
		if nodeAsn {
			okc++
			studyed[values[i]] = 1
		} else {
			errc++
		}
	}
	if errc == len(values) {
		return false
	}
	if okc == len(values) {
		delete(waitStudy, node) // 删除
		return true
	}
	return search(pMap, studyed, waitStudy, node)
}
