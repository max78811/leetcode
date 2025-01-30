package _71_JewelsandStones

// https://leetcode.com/problems/jewels-and-stones/description/
func numJewelsInStones(jewels string, stones string) int {
	m := make(map[string]struct{})
	for _, i := range jewels {
		m[string(i)] = struct{}{}
	}

	var counter int
	for _, i := range stones {
		_, exists := m[string(i)]
		if exists {
			counter++
		}
	}
	return counter
}
