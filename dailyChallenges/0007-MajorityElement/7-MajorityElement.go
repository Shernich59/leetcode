func majorityElement(nums []int) int {
	count, res := 0, 0

	for _, n := range nums {
		if count == 0 {
			res = n
		}
		if res == n {
			count += 1
		} else {
			count -= 1
		}
	}
	return res
}