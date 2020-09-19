package sorts

import "math/rand"

func generateRandonArray(n int) []int {
	nums := make([]int, 0)
	for len(nums) < n {
		nums = append(nums, rand.Intn(100))
	}
	return nums
}
