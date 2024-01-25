func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	arr := []rune(s)

	for l < r {
		left := unicode.ToLower(arr[l])
		right := unicode.ToLower(arr[r])

		if !isAlphaNum(left) {
			l++
			continue
		}

		if !isAlphaNum(right) {
			r--
			continue
		}

		if left != right {
			return false
		}
		l++
		r--
	}
	return true
}

func isAlphaNum(s rune) bool {
	return unicode.IsLetter(s) || unicode.IsDigit(s)
}
