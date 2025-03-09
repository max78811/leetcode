package RLEzip

import "strconv"

//AAABBCCCC
//A3B2C4

func RLEzip(input string) string {
	var result string
	var currentCount int
	var prevValue rune

	for index, letter := range input {
		if letter != prevValue || len(input)-1 == index {
			countStr := strconv.Itoa(currentCount)
			result += string(letter) + countStr
		}
		currentCount++
		prevValue = letter
	}
	return result
}
