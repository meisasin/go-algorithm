package main

import "fmt"

func main() {

	f := 7 / 0.5
	fmt.Println(f)

	f2 := 20 / 30.0
	fmt.Println(f2)
	f = 7 / f2
	fmt.Println(f)

}

/**
最低票价

在一个火车旅行很受欢迎的国度，你提前一年计划了一些火车旅行。在接下来的一年里，你要旅行的日子将以一个名为 days 的数组给出。每一项是一个从 1 到 365 的整数。
火车票有三种不同的销售方式：

- 一张为期一天的通行证售价为 costs[0] 美元；
- 一张为期七天的通行证售价为 costs[1] 美元；
- 一张为期三十天的通行证售价为 costs[2] 美元。
通行证允许数天无限制的旅行。 例如，如果我们在第 2 天获得一张为期 7 天的通行证，那么我们可以连着旅行 7 天：第 2 天、第 3 天、第 4 天、第 5 天、第 6 天、第 7 天和第 8 天。
返回你想要完成在给定的列表 days 中列出的每一天的旅行所需要的最低消费。


示例 1:
```
输入：days = [1,4,6,7,8,20], costs = [2,7,15]
输出：11
解释：
例如，这里有一种购买通行证的方法，可以让你完成你的旅行计划：
在第 1 天，你花了 costs[0] = $2 买了一张为期 1 天的通行证，它将在第 1 天生效。
在第 3 天，你花了 costs[1] = $7 买了一张为期 7 天的通行证，它将在第 3, 4, ..., 9 天生效。
在第 20 天，你花了 costs[0] = $2 买了一张为期 1 天的通行证，它将在第 20 天生效。
你总共花了 $11，并完成了你计划的每一天旅行。
```

示例 1:
```
输入：days = [1,2,3,4,5,6,7,8,9,10,30,31], costs = [2,7,15]
输出：17
解释：
例如，这里有一种购买通行证的方法，可以让你完成你的旅行计划：
在第 1 天，你花了 costs[2] = $15 买了一张为期 30 天的通行证，它将在第 1, 2, ..., 30 天生效。
在第 31 天，你花了 costs[0] = $2 买了一张为期 1 天的通行证，它将在第 31 天生效。
你总共花了 $17，并完成了你计划的每一天旅行。
```

提示：

1. 1 <= days.length <= 365
2. 1 <= days[i] <= 365
3. days 按顺序严格递增
4. costs.length == 3
5. 1 <= costs[i] <= 1000
```
*/

/**
太鸡难了，不由得有个疑问，动态规划真的科学吗？
TODO 后续再温习温习
*/
func mincostTickets(days []int, costs []int) int {
	memo := [366]int{}
	dayM := map[int]bool{}
	for _, d := range days {
		dayM[d] = true
	}

	var dp func(day int) int
	dp = func(day int) int {
		if day > 365 {
			return 0
		}
		if memo[day] > 0 {
			return memo[day]
		}
		if dayM[day] {
			memo[day] = min(min(dp(day+1)+costs[0], dp(day+7)+costs[1]), dp(day+30)+costs[2])
		} else {
			memo[day] = dp(day + 1)
		}
		return memo[day]
	}
	return dp(1)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
