package main

//https://leetcode.com/problems/two-sum/description/

func twoSum(nums []int, target int) []int {
	left := 0
	right := 1

	for left < right {
		if nums[left]+nums[right] == target {
			break
		}

		if len(nums)-1 <= right {
			left++
			right = left + 1
			continue
		}
		right++
	}

	return []int{left, right}
}
