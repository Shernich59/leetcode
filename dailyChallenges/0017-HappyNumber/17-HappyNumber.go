func isHappy(n int) bool {
	slow, fast := n, squareOfNum(n)

	for fast != 1 && slow != fast {
		slow = squareOfNum(slow)
		fast = squareOfNum(squareOfNum(fast))
	}
	return fast == 1
}

func squareOfNum(n int) int {
	output := 0

	for n > 0 {
		digit := n % 10
		output += digit * digit
		n /= 10
	}
	return output
}