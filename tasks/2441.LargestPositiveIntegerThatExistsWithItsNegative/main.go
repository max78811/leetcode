package main

//	I
//
// [-1,2,-3,3]
//
//	I
func findMaxK(nums []int) int {
	left := 0
	right := 1

	result := -1

	for {
		if left >= len(nums)-1 {
			break
		}
		if nums[left]*-1 == nums[right] || nums[left] == nums[right]*-1 {
			r := max(nums[left], nums[right])
			result = max(r, result)
		}
		if right >= len(nums)-1 {
			left++
			right = left + 1
			continue
		}
		right++
	}
	return result
}
