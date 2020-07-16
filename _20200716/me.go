package main

import "fmt"

/**
96. 不同的二叉搜索树
785. 判断二分图

给定一个无向图`graph`，当这个图为二分图时返回`true`。
如果我们能将一个图的节点集合分割成两个独立的子集A和B，并使图中的每一条边的两个节点一个来自A集合，一个来自B集合，我们就将这个图称为二分图。
`graph`将会以邻接表方式给出，`graph[i]`表示图中与节点`i`相连的所有节点。每个节点都是一个在`0`到`graph.length-1`之间的整数。这图中没有自环和平行边：
`graph[i]` 中不存在i，并且`graph[i]`中没有重复的值。

示例1：
```
示例 1:
输入: [[1,3], [0,2], [1,3], [0,2]]
输出: true
解释:
无向图如下:
0----1
|    |
|    |
3----2
我们可以将节点分成两组: {0, 2} 和 {1, 3}。
```


示例2：
```

示例 2:
输入: [[1,2,3], [0,2], [0,1,3], [0,2]]
输出: false
解释:
无向图如下:
0----1
| \  |
|  \ |
3----2
我们不能将节点分割成两个独立的子集。
```

注意:
- `graph 的长度范围为 [1, 100]。`
- `graph[i] 中的元素的范围为 [0, graph.length - 1]。`
- `graph[i] 不会包含 i 或者有重复的值。`
- `图是无向的: 如果j 在 graph[i]里边, 那么 i 也会在 graph[j]里边。`
*/

/**
模拟了半天才搞定，还好还好，下次再遇见就有思路发
*/
func IsBipartite(graph [][]int) bool {

	ok := make([]int, len(graph))
	ans := make([]int, len(graph))

	ans[0] = 1
	queue := []int{0}
	count := 0
	for len(queue) > 0 || count < len(graph) {
		if len(queue) == 0 {
			for i := 0; i < len(ok); i++ {
				if ok[i] == 0 {
					queue = append(queue, i)
					ans[i] = 1
					break
				}
			}
		}
		head := queue[0]
		queue = queue[1:]
		if ok[head] == 1 {
			continue
		}
		ok[head] = 1
		count++
		g := graph[head]
		for i := 0; i < len(g); i++ {
			if ans[g[i]] == ans[head] {
				return false
			}
			if ans[g[i]] == 0 {
				ans[g[i]] = ans[head] * -1
			}
			queue = append(queue, g[i])
		}
	}
	// fmt.Println(ok)
	fmt.Println(ans)
	return true
}
