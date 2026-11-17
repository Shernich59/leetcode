func containsNearbyDuplicate(nums []int, k int) bool {
	window := make(map[int]struct{})
	L := 0

	for R := 0; R < len(nums); R++ {
		if R-L > k {
			delete(window, nums[L])
			L++
		}
		if _, exists := window[nums[R]]; exists {
			return true
		}
		window[nums[R]] = struct{}{}
	}
	return false
}