func majorityElement(nums []int) int {
    count := make(map[int]int)
    result, maxCount := 0, 0

    for _, n := range nums {
        count[n]++

        if count[n] > maxCount {
            result = n
            maxCount = count[n]
        }
    }

    return result
}