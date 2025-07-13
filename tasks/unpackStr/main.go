package main

import (
	"fmt"
	"strconv"
)

func unpackStr(s string) string {
	if s == "" {
		return ""
	}

	var result string
	var prev string
	var prevIsNum bool

	r := []rune(s)

	for i, v := range r {
		num, err := strconv.Atoi(string(v))
		if err == nil {
			if prevIsNum || i == 0 {
				return ""
			}

			// is num
			for j := 0; j < num; j++ {
				result += prev
			}
			prevIsNum = true
			continue
		}
		result += string(v)
		prev = string(v)
	}
	return result
}

func main() {
	fmt.Println(unpackStr("d\\n5abc"))
}
