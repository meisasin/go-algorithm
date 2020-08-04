package main

import "strconv"

/**
114. 二叉树展开为链表

给定一个二叉树，原地将它展开为一个单链表。

示例1：
```
    1
   / \
  2   5
 / \   \
3   4   6
```

将其展开为：
```
1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6
```
*/

/**
哈哈，有点蠢。话说效率咋就这么低呢
*/
func AddStrings(num1 string, num2 string) string {

	carryBit := 0
	maxLen := max(len(num1), len(num2))
	ans := ""
	for i := 0; i < maxLen; i++ {
		count := carryBit
		if i < len(num1) {
			count += int(num1[len(num1)-i-1] - '0')
		}
		if i < len(num2) {
			count += int(num2[len(num2)-i-1] - '0')
		}
		ans = strconv.Itoa(count%10) + ans
		carryBit = count / 10
	}
	if carryBit == 0 {
		return ans
	}
	return "1" + ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
