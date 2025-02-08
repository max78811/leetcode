package _42_ValidAnagram

// https://leetcode.com/problems/valid-anagram/description/?envType=problem-list-v2&envId=string
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[string]int) //val count
	for _, i := range s {
		m[string(i)]++
	}

	for _, i := range t {
		val, exist := m[string(i)]
		if !exist {
			return false
		}
		if val <= 0 {
			return false
		}
		m[string(i)]--
	}
	return true
}
