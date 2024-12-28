package main

//               I
// ["h","e","l","l","o","y"]
//               I

// y o e l l h

//           I
// ["h","e","l","l","o"]
//           I

// o l l e h

//https://leetcode.com/problems/reverse-string/description/

func reverseString(s []byte) {
	left := 0
	right := len(s) - 1

	for left <= right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
