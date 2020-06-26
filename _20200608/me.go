package main

/**
等式方程的可满足性

给定一个由表示变量之间关系的字符串方程组成的数组，每个字符串方程 `equations[i]` 的长度为 4，并采用两种不同的形式之一："a==b" 或 "a!=b"。在这里，a 和 b 是小写字母（不一定不同），表示单字母变量名。
只有当可以将整数分配给变量名，以便满足所有给定的方程时才返回 true，否则返回 false。

示例1：
```
输入：["a==b","b!=a"]
输出：false
解释：如果我们指定，a = 1 且 b = 1，那么可以满足第一个方程，但无法满足第二个方程。没有办法分配变量同时满足这两个方程。
```

示例2：
```
输出：["b==a","a==b"]
输入：true
解释：我们可以指定 a = 1 且 b = 1 以满足满足这两个方程。
```

示例3：
```
输入：["a==b","b==c","a==c"]
输出：true
```

示例4：
```
输入：["a==b","b!=c","c==a"]
输出：false
```

示例5：
```
输入：["c==c","b==d","x!=z"]
输出：true
```
*/

/**
没做出来 ಥ_ಥ
*/
func EquationsPossible(equations []string) bool {
	ans := true
	// fmt.Println(equations)

	emap := make(map[byte]int)
	// fmt.Println(emap)

	count := 0
	for _, equ := range equations {
		if equ[1] == '=' {
			val1, ok1 := emap[equ[0]]
			val2, ok2 := emap[equ[3]]
			if ok1 && ok2 {
				ans = ans && val1 == val2
			} else if ok1 {

				emap[equ[3]] = val1
			} else if ok2 {
				emap[equ[0]] = val2
			} else {
				emap[equ[0]] = count
				emap[equ[3]] = count
				count++
			}
		}
		if !ans {
			return false
		}
	}

	for _, equ := range equations {
		if equ[1] == '!' {
			if equ[0] == equ[3] {
				return false
			}
			val1, ok1 := emap[equ[0]]
			val2, ok2 := emap[equ[3]]
			if ok1 && ok2 {
				ans = ans && val1 != val2
			}
		}
		if !ans {
			return false
		}
	}
	return ans
}
