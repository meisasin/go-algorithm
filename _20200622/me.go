package main

import (
	"fmt"
	"strings"
)

/**
面试题 16.18. 模式匹配

你有两个字符串，即`pattern`和`value`。
`pattern`字符串由字母`"a"`和`"b"`组成，用于描述字符串中的模式。
例如，字符串`"catcatgocatgo"`匹配模式`"aabab"`（其中`"cat"`是`"a"`，`"go"`是`"b"`），该字符串也匹配像`"a"`、`"ab"`和`"b"`这样的模式。
但需注意`"a"`和`"b"`不能同时表示相同的字符串。编写一个方法判断`value`字符串是否匹配`pattern`字符串。

示例1：
```
输入： pattern = "abba", value = "dogcatcatdog"
输出： true
```

示例2：
```
输入： pattern = "abba", value = "dogcatcatfish"
输出： false
```

示例3：
```
输入： pattern = "aaaa", value = "dogcatcatdog"
输出： false
```

示例4：
```
输入： pattern = "abba", value = "dogdogdogdog"
输出： true
解释： "a"="dogdog",b=""，反之也符合规则
```

提示：

- `0 <= len(pattern) <= 1000`
- `0 <= len(value) <= 1000`
- 你可以假设`pattern`只包含字母`"a"`和`"b"`，`value`仅包含小写字母。
*/

/**
简单题，代码写的确实有点长了，一点都不优雅，一会去找找 Go 自带的校验和转换方法，就不用自己写方法了
*/

/**
修修补补，思路到底还是错的
ERROR
*/
func PatternMatching(pattern string, value string) bool {
	vlen := len(value)
	plen := len(pattern)

	if plen == 0 && vlen == 0 {
		return true
	}
	if plen == 0 {
		return false
	}

	if vlen == 0 {
		return (len(strings.ReplaceAll(pattern, "a", "")) == 0) || (len(strings.ReplaceAll(pattern, "b", "")) == 0)
	}

	for i := 0; i < vlen; i++ {
		// a
		a := value[0:i]

		step := i
		for step < vlen {
			// b
			b := value[i:step]
			fmt.Println("a is ", a, ", b is ", b)
			// 如果能被整除，说明数量能对得上
			if len(a)+len(b) > 0 {
				newStr := strings.ReplaceAll(strings.ReplaceAll(pattern, "a", a), "b", b)
				fmt.Println("a is ", a, ", b is ", b, ", String is "+value+" - "+newStr)
				if newStr == value {
					return true
				}
			}
			step++
		}
	}
	return false
}
