package other

import (
	"fmt"
	"testing"
)

func TestModifyString(t *testing.T) {
	fmt.Println(modifyString("?zs"))
	fmt.Println(modifyString("j?qg??b")) // "jaqgacb"
}

func modifyString(s string) string {

	sl := len(s)
	var ans []byte
	if sl == 1 {
		if s[0] == '?' {
			return "a"
		}
		return s
	}
	s = "-" + s + "-"
	sl = len(s)
	for i := 1; i < sl-1; i++ {
		if s[i] != '?' {
			ans = append(ans, s[i])
		} else {
			for j := 'a'; j <= 'z'; j++ {
				if uint8(j) != s[i-1] && uint8(j) != s[i+1] {
					if len(ans) > 0 && ans[len(ans)-1] == uint8(j) {
						continue
					}
					ans = append(ans, uint8(j))
					break
				}
			}
		}
	}
	fmt.Println(ans)
	return string(ans)

}

func TestNumTriplets(t *testing.T) {
	//fmt.Println(numTriplets([]int { 7,4}, []int { 5,2,8,9 } ))
	fmt.Println(numTriplets([]int{}, []int{}))
}

func numTriplets(nums1 []int, nums2 []int) int {

	n1l := len(nums1)
	n2l := len(nums2)

	ans := 0

	n1map := make(map[int][]int)
	for i := 0; i < n1l; i++ {
		n1map[nums1[i]] = append(n1map[nums1[i]], i)
	}
	n2map := make(map[int][]int)
	for i := 0; i < n2l; i++ {
		n2map[nums2[i]] = append(n2map[nums2[i]], i)
	}

	for i := 0; i < n1l; i++ {
		count := nums1[i] * nums1[i]
		for j := 0; j < n2l; j++ {
			kint := count / nums2[j]
			if kint*nums2[j] != count {
				continue
			}
			idxs := n2map[kint]
			if len(idxs) > 0 {
				ding := len(idxs)
				for x := 0; x < len(idxs); x++ {
					if idxs[x] > j {
						ding = x
						break
					}
				}
				ans += len(idxs) - ding
			}
		}
	}
	for i := 0; i < n2l; i++ {
		count := nums2[i] * nums2[i]
		for j := 0; j < n1l; j++ {
			kint := count / nums1[j]
			if kint*nums1[j] != count {
				continue
			}
			idxs := n1map[kint]
			if len(idxs) > 0 {
				ding := len(idxs)
				for x := 0; x < len(idxs); x++ {
					if idxs[x] > j {
						ding = x
						break
					}
				}
				ans += len(idxs) - ding
			}
		}
	}
	fmt.Println("-----------")
	return ans
}

func TestBBB(t *testing.T) {

	a := "abcd"
	b := 2e9

	fmt.Println(a)
	fmt.Println(b)
}
