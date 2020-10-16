package main

/**
977. 有序数组的平方

给定一个按非递减顺序排序的整数数组 A，返回每个数字的平方组成的新数组，要求也按非递减顺序排序。

示例1：
```
输入：[-4,-1,0,3,10]
输出：[0,1,9,16,100]
```

示例2：
```
输入：[-7,-3,2,3,11]
输出：[4,9,9,49,121]
```

提示：
- 1 <= A.length <= 10000
- -10000 <= A[i] <= 10000
- A 已按非递减顺序排序。
*/

// ------------------------------------------------------------------------------------------

/**
...
*/

// ------------------------------------------------------------------------------------------

func SortedSquares(A []int) []int {

	var stack []int

	point := 0
	for i := 0; i < len(A); i++ {

		if A[i] < 0 {
			stack = append(stack, A[i]*A[i])
		} else {
			res := A[i] * A[i]
			for len(stack) > 0 && stack[len(stack)-1] < res {
				A[point] = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				point++
			}
			A[point] = res
			point++
		}
	}
	for i := len(stack) - 1; i >= 0; i-- {
		A[point] = stack[i]
		point++
	}
	return A

}

//func SortedSquares(A []int) []int {
//
//	for i := 0 ; i < len(A) ; i ++ {
//		if A[i] >= 0 {
//			break
//		}
//		A[i] = A[i] * -1
//	}
//	sort.Ints(A)
//	for i := 0 ; i < len(A) ; i ++ {
//		A[i] = A[i] * A[i]
//	}
//	return A
//}
