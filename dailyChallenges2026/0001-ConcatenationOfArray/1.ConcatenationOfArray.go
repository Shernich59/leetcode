func getConcatenation(nums []int) []int {
    result := []int{}

    for i:=0; i<2; i++ {
        for _, n := range nums {
            result = append(result, n)
        }
    }
    return result
}