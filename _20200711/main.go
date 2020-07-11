package main

import "sort"

func main() {

}

var a, c []int

func countSmaller(nums []int) []int {
	resultList := []int{}
	discretization(nums)
	c = make([]int, len(nums)+5)
	for i := len(nums) - 1; i >= 0; i-- {
		id := getId(nums[i])
		resultList = append(resultList, query(id-1))
		update(id)
	}
	for i := 0; i < len(resultList)/2; i++ {
		resultList[i], resultList[len(resultList)-1-i] = resultList[len(resultList)-1-i], resultList[i]
	}
	return resultList
}

func lowBit(x int) int {
	return x & (-x)
}

func update(pos int) {
	for pos < len(c) {
		c[pos]++
		pos += lowBit(pos)
	}
}

func query(pos int) int {
	ret := 0
	for pos > 0 {
		ret += c[pos]
		pos -= lowBit(pos)
	}
	return ret
}

func discretization(nums []int) {
	set := map[int]struct{}{}
	for _, num := range nums {
		set[num] = struct{}{}
	}
	a = make([]int, 0, len(nums))
	for num := range set {
		a = append(a, num)
	}
	sort.Ints(a)
}

func getId(x int) int {
	return sort.SearchInts(a, x) + 1
}
