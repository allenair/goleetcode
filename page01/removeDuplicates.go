package page01

import "fmt"

func RemoveDuplicates(nums []int) int {
	lastNum, num := 0, 0
	for i := 0; i < len(nums); i++ {
		if i > 0 && lastNum == nums[i] {
			for k := i; k < len(nums); k++ {
				if nums[k] != lastNum {
					lastNum = nums[k]
					nums = append(nums[:i], nums[k:]...)
					num++
					break
				} else if k == len(nums)-1 {
					nums = nums[:i]
				}
			}
		} else {
			lastNum = nums[i]
			num++
		}

	}
	fmt.Println(nums)
	return num
}
