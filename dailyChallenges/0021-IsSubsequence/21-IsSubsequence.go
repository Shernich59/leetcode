func isSubsequence(s string, t string) bool {

	i, j := 0, 0

	if len(s) > len(t) {
		return false
	}

	for i < len(s) {
		for j < len(t) && s[i] != t[j] {
			j++
		}
		if j == len(t) {
			return false
		}
		i++
		j++
	}
	return true
}