package main

func canConstruct(ransomNote string, magazine string) bool {
	n := make(map[string]int) //var count
	for _, j := range magazine {
		n[string(j)]++
	}

	h := []rune(ransomNote)
	for _, g := range h {
		strG := string(g)
		counter, exist := n[strG]
		if exist {
			if counter <= 0 {
				return false
			}
			n[strG]--
		} else {
			return false
		}
	}

	return true
}
