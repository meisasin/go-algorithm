package other

/**
01背包问题，救黄金最优解
@param w 工人数量
@param p 表示第 i 个金矿开采所需要的工人数量
@param g 表示第 i 个金矿储量
F(n, w)
*/
func Package_01(w int, p []int, g []int) int {

	//res := make([][]int, len(g) + 1)
	//for i := range res { res[i] = make([]int, w + 1)}
	res := make([]int, w+1)
	for i := 0; i < len(p); i++ { // 第 n 个 金矿
		for j := len(res) - 1; j >= p[i]; j-- { // 当有 w 个工人时
			res[j] = max(res[j], res[j-p[i]]+g[i])
		}
	}
	return res[w]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
