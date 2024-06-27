package easy

import "slices"

func removeElement(nums []int, val int) int {
	nums = slices.DeleteFunc(nums, func(i int) bool {
		return val == i
	})
	return len(nums)
}
