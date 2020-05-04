package page01

import "math"

func LongestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	res := make([]rune, 0, 10)
	allStr := make(map[int][]rune)
	minLen := math.MaxInt32
	for index, str := range strs {
		allStr[index] = []rune(str)
		if len(str) < minLen {
			minLen = len(str)
		}
	}
	for i := 0; i < minLen; i++ {
		c := allStr[0][i]
		for _, v := range allStr {
			if c != v[i] {
				return string(res)
			}
		}
		res = append(res, c)
	}
	return string(res)
}
