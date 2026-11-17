func romanToInt(s string) int {
	roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	res := 0

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && roman[string(s[i])] < roman[string(s[i+1])] {
			res -= roman[string(s[i])]
		} else {
			res += roman[string(s[i])]
		}
	}

	return res
}
