// In golang, strings are immutable
// we are not allowed to change the individual characters of a string directly.
// Using bytes allows us to work with mutable data : )
func isIsomorphic(s string, t string) bool {
	mapST := make(map[byte]byte)
	mapTS := make(map[byte]byte)

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		c1 := s[i]
		c2 := t[i]

		if (mapST[c1] != 0 && mapST[c1] != c2) || (mapTS[c2] != 0 && mapTS[c2] != c1) {
			return false
		} else {
			mapST[c1] = c2
			mapTS[c2] = c1
		}
	}
	return true
}