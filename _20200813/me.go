package main

import (
	"strconv"
)

/**
43. 字符串相乘

给定两个以字符串形式表示的非负整数 `num1` 和 `num2`，返回 `num1` 和 `num2` 的乘积，它们的乘积也表示为字符串形式。


示例1：
```
输入: num1 = "2", num2 = "3"
输出: "6"
```

示例2：
```
输入: num1 = "123", num2 = "456"
输出: "56088"
```

说明：
- `num1 和 num2 的长度小于110。`
- `num1 和 num2 只包含数字 0-9。`
- `num1 和 num2 均不以零开头，除非是数字 0 本身。`
- `不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。`
*/

func Multiply(num1 string, num2 string) string {

	if num1 == "" || num2 == "" {
		return ""
	}
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	il, jl := len(num1), len(num2)
	res := make([][]int, il)

	for i := 0; i < il; i++ {
		res[i] = make([]int, jl+1)
	}
	for i := il - 1; i >= 0; i-- {
		a, _ := strconv.Atoi(string(num1[i]))
		for j := jl - 1; j >= 0; j-- {
			b, _ := strconv.Atoi(string(num2[j]))
			value := a * b
			res[il-i-1][jl-j-1] = res[il-i-1][jl-j-1] + value%10
			res[il-i-1][jl-j] = value / 10
		}
	}

	ans := ""
	jin := 0
	for i := 0; i < il; i++ {
		currPos, offsit, currVal := i, 0, jin
		for currPos >= 0 && offsit <= jl {
			currVal += res[currPos][offsit]
			currPos--
			offsit++
		}
		jin = currVal / 10
		ans = strconv.Itoa(currVal%10) + ans
	}

	for i := 1; i < jl+1; i++ {
		currPos, offsit, currVal := il-1, i, jin
		for currPos >= 0 && offsit <= jl {
			currVal += res[currPos][offsit]
			currPos--
			offsit++
		}
		jin = currVal / 10
		if i == jl && currVal == 0 {
			break
		}
		ans = strconv.Itoa(currVal%10) + ans
	}

	//fmt.Println(res)
	return ans
}
