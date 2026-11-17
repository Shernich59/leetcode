func lengthOfLongestSubstring(s string) int {
	hashSet := make(map[byte]bool)
	l, res := 0, 0

	for r, _ := range s {
		for hashSet[s[r]] {
			delete(hashSet, s[l])
			l++
		}
		hashSet[s[r]] = true
		res = max(res, r-l+1)
	}
	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
