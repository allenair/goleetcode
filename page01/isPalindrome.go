package page01

import "strconv"

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	arr := []byte(strconv.Itoa(x))
	for i, k := 0, len(arr)-1; i < k; {
		if arr[i] != arr[k] {
			return false
		}
		i++
		k--
	}
	return true
}
