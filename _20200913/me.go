package main

/**
79. 单词搜索

给定一个二维网格和一个单词，找出该单词是否存在于网格中。
单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

示例1:
```
board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

给定 word = "ABCCED", 返回 true
给定 word = "SEE", 返回 true
给定 word = "ABCB", 返回 false
```

提示：
- `board 和 word 中只包含大写和小写英文字母。`
- `1 <= board.length <= 200`
- `1 <= board[i].length <= 200`
- `1 <= word.length <= 10^3`
*/

/**
写的是啥自己也不知道，居然一把过了
*/
var rnp = []int{-1, 1, 0, 0}
var cnp = []int{0, 0, -1, 1}

func w(board [][]byte, word string) bool {

	rl, cl := len(board), len(board[0])
	wl := len(word)
	for i := 0; i < rl; i++ {
		for j := 0; j < cl; j++ {
			if dfs(board, []int{i, j}, word, wl, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, posi []int, word string, wl, point int) bool {
	rp, cp := posi[0], posi[1]
	if board[rp][cp] != word[point] {
		return false
	}
	if point == wl-1 {
		return true
	}
	tmp := board[rp][cp]
	board[rp][cp] = '0'
	for i := 0; i < 4; i++ {
		if rp+rnp[i] >= 0 && rp+rnp[i] < len(board) && cp+cnp[i] >= 0 && cp+cnp[i] < len(board[0]) {
			if dfs(board, []int{rp + rnp[i], cp + cnp[i]}, word, wl, point+1) {
				return true
			}
		}
	}

	board[rp][cp] = tmp
	return false
}
