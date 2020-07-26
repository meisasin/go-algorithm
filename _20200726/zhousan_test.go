package main

import (
	"fmt"
	"testing"
)

func TestSimple(t *testing.T) {

	fmt.Println(restoreString("abc", []int{0, 1, 2}))
	fmt.Println(restoreString("aaiougrt", []int{4, 0, 2, 6, 7, 3, 1, 5}))

}

func restoreString(s string, indices []int) string {

	ans := make([]uint8, len(indices))

	for i := 0; i < len(indices); i++ {
		ans[indices[i]] = s[i]
	}
	return string(ans)
}

func minFlips(target string) int {
	count := 0
	for i := 0; i < len(target); i++ {
		if target[i] == '0' {
			if count%2 == 1 {
				count++
			}
		} else {
			if count%2 == 0 {
				count++
			}
		}
	}
	return count
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countPairs(root *TreeNode, distance int) int {

	ans, count := count(root, distance)
	fmt.Println(ans)
	return count
}

func count(root *TreeNode, distcace int) ([]int, int) {
	if root.Left == nil && root.Right == nil {
		return []int{1}, 0
	}
	var lTmp []int
	var rTmp []int
	ccount := 0
	if root.Left != nil {
		tmp, cc := count(root.Left, distcace)
		ccount += cc
		lTmp = tmp
	}
	if root.Right != nil {
		tmp, cc := count(root.Right, distcace)
		ccount += cc
		rTmp = tmp
	}

	if root.Left != nil && root.Right != nil {
		for i := 0; i < len(lTmp); i++ {
			for j := 0; j < len(rTmp); j++ {
				if lTmp[i]+rTmp[j] <= distcace {
					ccount++
				}
			}
		}
	}
	var ans []int
	ans = append(ans, lTmp...)
	ans = append(ans, rTmp...)
	for i := 0; i < len(ans); i++ {
		ans[i] = ans[i] + 1
	}
	return ans, ccount
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestA(t *testing.T) {
	fmt.Println(minFlips("10111"))
	fmt.Println(minFlips("101"))
	fmt.Println(minFlips("00000"))
	fmt.Println(minFlips("001011101"))
}
