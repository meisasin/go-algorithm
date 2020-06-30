package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestFindKthLargest(t *testing.T) {

	fmt.Println(1 / 2)
}

func TestMySort(t *testing.T) {
	arr := []int{1, 4, 9, 56, 11, 2, 2, 3, 7, 21, 78, 102, 8, 99, 7}
	MySort(arr)

	fmt.Println(arr)
}

func MySort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

func TestQuickSort(t *testing.T) {

}

func TestComp(t *testing.T) {
	// [1 2 2 3 4 7 7 8 9 11 21 56 78 99 102]
	//arr := []int{1, 4, 9, 56, 11, 2, 2, 3, 7, 21, 78, 102, 8, 99, 7 }
	le := 10000
	arr := make([]int, le)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < le; i++ {
		arr[i] = rand.Intn(le)
	}
	timeComp("快排", arr, QuickSort)
	timeComp("冒泡", arr, MySort)
}
func timeComp(message string, arr []int, comp func([]int)) {
	start := time.Now()
	comp(arr)
	end := time.Now()
	fmt.Println("[", message, "]Start:", start, ", End:", end, " --> ", end.Sub(start).Microseconds())
}

func TestRand(t *testing.T) {

	rand.Seed(time.Now().UnixNano())
	//for  {
	//	//fmt.Println(rand.Int31n(10))
	//	fmt.Println(rand.Int())
	//	time.Sleep(time.Second * 1)
	//}
}

/**
快速排序
*/
func QuickSort(nums []int) {
	Quick(nums, 0, len(nums)-1)
}

func Quick(nums []int, left int, right int) {
	if left >= right {
		return
	}

	rand.Seed(time.Now().UnixNano())
	baseIdx := left + rand.Intn(right-left+1)
	base := nums[baseIdx]
	var minS []int
	var maxS []int
	var midS []int
	for i := left; i <= right; i++ {
		if i == baseIdx {
			continue
		}

		if nums[i] > base {
			maxS = append(maxS, nums[i])
		} else if nums[i] < base {
			minS = append(minS, nums[i])
		} else {
			midS = append(midS, nums[i])
		}
	}
	l_r := left + len(minS)
	r_l := left + len(minS) + len(midS)
	minS = append(minS, midS...)
	minS = append(minS, maxS...)
	for i, j := left, 0; i <= right; {
		nums[i] = minS[j]
		i++
		j++
	}
	Quick(nums, left, l_r)
	Quick(nums, r_l, right)
}
