package other

func TwoSum(nums []int, target int) []int {
	var aMap = make(map[int]int)

	for i := 0; i < len(nums); i++ {
		idx, ok := aMap[target-nums[i]]
		if ok {
			return []int{idx, i}
		}
		aMap[nums[i]] = i
	}
	return []int{-1, -1}
}
