package main

func main() {

}

func smallerNumbersThanCurrent(nums []int) []int {
	cnt := [101]int{}
	for _, v := range nums {
		cnt[v]++
	}
	for i := 0; i < 100; i++ {
		cnt[i+1] += cnt[i]
	}
	ans := make([]int, len(nums))
	for i, v := range nums {
		if v > 0 {
			ans[i] = cnt[v-1]
		}
	}
	return ans
}
