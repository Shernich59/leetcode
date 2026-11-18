class Solution:
    def majorityElement(self, nums: list[int]) -> int:
        count = {}
        result, maxCount = 0, 0

        for n in nums:
            count[n] = 1 + count.get(n, 0)
            result = n if count[n] > maxCount else result
            maxCount = max(count[n], maxCount)
        return result