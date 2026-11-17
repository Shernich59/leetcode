func canConstruct(ransomNote string, magazine string) bool {
	//rune represent int32 type (Unicode character)
	_mag := make(map[rune]int)

	// Populate the _mag with character counts from magazine
	for _, char := range magazine {
		_mag[char]++
	}

	// Check if the characters in ransomNote can be constructed from magazine
	for _, char := range ransomNote {
		count, found := _mag[char]
		if !found || count == 0 {
			return false
		}
		_mag[char]--
	}
	return true
}
