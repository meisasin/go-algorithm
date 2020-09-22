package main

/**
968. 监控二叉树

给定一个二叉树，我们在树的节点上安装摄像头。
节点上的每个摄影头都可以监视其父对象、自身及其直接子对象。
计算监控树的所有节点所需的最小摄像头数量。


示例 1：
![_1.png](./source/_1.png)
```
输入：[0,0,null,0,0]
输出：1
解释：如图所示，一台摄像头足以监控所有节点。
```


示例 2：
![_2.png](./source/_2.png)
```
输入：[0,0,null,0,null,0,null,null,0]
输出：2
解释：需要至少两个摄像头来监视树的所有节点。 上图显示了摄像头放置的有效位置之一。
```

提示：
- `给定树的节点数的范围是 [1, 1000]。`
- `每个节点的值都是 0。`
*/

/**
困难题日常用时一下午做不出来
error
*/
func MinCameraCover(root *TreeNode) int {

	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	ok, unok, half := dfs(root)
	return min(min(ok, unok), half)
}

// @return !1 > 选当前节点时最小值，!2 > 不选当前节点时最小值，!3 > 不选当前节点时，且不选下级节点时的最小值，
func dfs(root *TreeNode) (int, int, int) {
	if root == nil {
		return 0, 0, 1
	}

	lok, lunok, lhalf := dfs(root.Left)
	rok, runok, rhalf := dfs(root.Right)

	cok := 1 + min(lunok, lhalf) + min(runok, rhalf)
	cunok := lok + rok
	chalf := lunok + runok
	return cok, cunok, chalf
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
