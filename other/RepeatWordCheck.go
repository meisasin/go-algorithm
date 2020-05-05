package other

/**
判断两个字符串是否重复
*/

func CheckNoRepeat(s1 string, s2 string) bool {

	if s1 == "" || s2 == "" {
		return true
	}

	bitmark1, bitmark2 := 0, 0
	for _, i := range s1 {
		bitmark1 |= 1 << (i - 'a')
	}
	for _, i := range s2 {
		bitmark2 |= 1 << (i - 'a')
	}
	return bitmark1&bitmark2 == 0
}
