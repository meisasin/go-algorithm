package main

/**
面试题64. 求1+2+…+n

求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。


示例1：
```
输入: n = 3
输出: 6
```

示例2：
```
输入: n = 9
输出: 45
```

限制：

- `1 <= n <= 10000`
*/

/**
一直想着递归时返回和，绕进去了没出来，其实还真是，在递归函数里面计算，在外部用个全局变量累加就可以了
*/
func sumNumsWithMe(n int) int {

	ans := 0
	var sumR func(int) bool
	sumR = func(nu int) bool {
		ans += nu
		nu--
		return nu > 0 && sumR(nu)
	}
	sumR(n)
	return ans
}
