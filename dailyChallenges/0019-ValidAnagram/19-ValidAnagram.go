func isAnagram(s string, t string) bool {
	sMap := make(map[rune]int)
	tMap := make(map[rune]int)

	if len(s) != len(t) {
		return false
	}

	for _, char := range s {
		sMap[char]++
	}

	for _, char := range t {
		tMap[char]++
	}

	for key, value1 := range sMap {
		if value2, ok := tMap[key]; !ok || value1 != value2 {
			return false
		}
	}
	return true
}