package main

//   I
//[1,2,2,1] [2,2]
//           I

//https://leetcode.com/problems/intersection-of-two-arrays/description/

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]struct{})

	left := 0
	right := 0
	c := 0

	for {
		if nums1[left] == nums2[right] {
			_, exist := m[nums1[left]]
			if !exist {
				m[nums1[left]] = struct{}{}
				c++
			}
		}

		if left >= len(nums1)-1 {
			right++
			if right > len(nums2)-1 {
				break
			}
			left = 0
			continue
		}
		left++
	}

	result := make([]int, 0, c)
	for k, _ := range m {
		result = append(result, k)
	}
	return result
}
