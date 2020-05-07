package other

import "fmt"

/**
@desc
	KMP 字符串匹配算法
*/

/**
计算字符串 s 对应的前缀表
计算规则：
	例： ababc
	分解： 	a  					->  前后无一致的为 0
		  	a b 				->  前后无一致的为 0
			a	b	a			->  第一位和最后一位一致为 1
			a	b	a	b 		->  前两位和最后两位一致为 2
			a	b	a	b	c	->  前后无一致的为 0
	答案：[0, 0, 1, 2, 0] 为了方便进行 KMP 计算，将数组往右移一位，最左位赋值为 -1 即变成 [-1, 0, 0, 1, 2]
	提示：可以使用动态规划
*/
func PrefixTable(s string) []int {

	var prefixTable []int = make([]int, len(s)-1)
	prefixTable = append(prefixTable, 1)

	// 开始只有一位，所以为0
	prefixTable[0] = 0
	// 从0开始匹配，如果匹配到一位就加 1，如果没匹配到，说明
	count := 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[count] {
			count++
		} else if s[i] == s[0] {
			count = 1
		} else {
			count = 0
		}
		prefixTable[i] = count
	}
	return MovePrefixTable(prefixTable)
}

/**
为了方便进行 KMP 计算，将数组往右移一位，最左位赋值为 -1
例：[0, 0, 1, 2, 0] 即变成 [-1, 0, 0, 1, 2]
*/
func MovePrefixTable(prefixTable []int) []int {

	for i := len(prefixTable) - 1; i > 0; i-- {
		prefixTable[i] = prefixTable[i-1]
	}
	prefixTable[0] = -1
	return prefixTable
}

/**
KMP 匹配在字符串 text中的 pattern 的索引
*/
func KMPSearch(text string, pattern string) []int {

	tp := 0
	pp := 0

	prefixTable := PrefixTable(pattern)
	// 用来存储匹配到的数组角标
	var idxArr []int

	for len(text)-tp-(len(pattern)-pp) >= 0 {

		fmt.Printf("TP: %d, PP: %d, len(idxArr): %d\n", tp, pp, len(idxArr))
		// 如果值相同，text 与 pattern 的索引同时往前走
		if text[tp] == pattern[pp] {
			tp++
			pp++
		} else {
			// 如果值不同，tp 不动，pp 退回到 prefixTable 对应的角标
			pp = prefixTable[pp]
		}

		// 如果 索引退回到 -1 那么 tp 需要向前走
		if pp == -1 {
			tp++
			pp = 0

		} else if pp == len(pattern) {
			// 如果匹配的数据已经走到头，说明全匹配了，将匹配到的开始索引放到数组角标，pp 置为 0
			idxArr = append(idxArr, tp-pp)
			// 将 tp 重置为 找到的 idx 重新匹配新的
			tp = tp - pp + 1
			pp = 0
		}

	}

	return idxArr

}

func GetNext(s string) []int {
	next := make([]int, len(s)+1)
	next[0] = -1
	k := -1
	j := 0
	for j < len(s)-1 {
		//fmt.Println("k=",k,"j=",j, "next =",next)
		//这里,k表示next[j-1],且s[k]表示前缀,s[j]表示后缀
		//注:k==-1表示未找到k前缀与k后缀相等,首次分析可先忽略
		if k == -1 || s[j] == s[k] {
			j++
			k++
			next[j+1] = k
		} else { //s[j]与s[k]不匹配,继续递归计算前缀s[next[k]]
			k = next[k]
		}
	}

	return next
}

func Kmp(s, p string, next []int) int {
	i, j := 0, 0
	for i < len(s) && j < len(p) {
		if j == -1 || s[i] == p[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == len(p) {
		return i - j
	}
	return -1
}
