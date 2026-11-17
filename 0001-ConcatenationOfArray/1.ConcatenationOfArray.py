class Solution:
    def getConcatenation(self, nums: list[int]) -> list[int]:
        result = []

        for i in range(2):
            for n in nums:
                result.append(n)
        return result