package main

/**
925. 长按键入

你的朋友正在使用键盘输入他的名字 name。偶尔，在键入字符 c 时，按键可能会被长按，而字符可能被输入 1 次或多次。
你将会检查键盘输入的字符 typed。如果它对应的可能是你的朋友的名字（其中一些字符可能被长按），那么就返回 True。

示例1：
```
输入：name = "alex", typed = "aaleex"
输出：true
解释：'alex' 中的 'a' 和 'e' 被长按。
```

示例2：
```
输入：name = "saeed", typed = "ssaaedd"
输出：false
解释：'e' 一定需要被键入两次，但在 typed 的输出中不是这样。
```

示例3：
```
输入：name = "leelee", typed = "lleeelee"
输出：true
```

示例4：
```
输入：name = "laiden", typed = "laiden"
输出：true
解释：长按名字中的字符并不是必要的。
```

提示：
- name.length <= 1000
- typed.length <= 1000
- name 和 typed 的字符都是小写字母。
*/

// ------------------------------------------------------------------------------------------

/**
优化了一下代码，这就简洁多了
*/

// ------------------------------------------------------------------------------------------
func IsLongPressedName(name string, typed string) bool {

	nl, tl := len(name), len(typed)
	np, tp := 0, 0
	for np < nl && tp < tl && name[np] == typed[tp] {
		nv, tv := name[np], typed[tp]
		nc, tc := 0, 0
		for np < nl && name[np] == nv {
			nc++
			np++
		}
		for tp < tl && typed[tp] == tv {
			tc++
			tp++
		}
		if nc > tc {
			return false
		}
	}
	return np == nl && tp == tl
}
