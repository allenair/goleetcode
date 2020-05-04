package page01

func TwoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	for i, v := range nums {
		if k, ok := dict[target-v]; ok {
			return []int{i, k}
		} else {
			dict[v] = i
		}
	}

	return nil
}
