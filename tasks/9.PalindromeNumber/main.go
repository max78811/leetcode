package main

import (
	"strconv"
)

//https://leetcode.com/problems/palindrome-number/description/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	strX := strconv.Itoa(x)
	r := []rune(strX)

	reversed := make([]rune, 0, len(r))
	for i := len(r) - 1; i >= 0; i-- {
		reversed = append(reversed, r[i])
	}

	if string(reversed) == string(r) {
		return true
	}

	return false
}
