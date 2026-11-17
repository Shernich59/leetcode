class Solution(object):
    def isHappy(self, n):
        """
        :type n: int
        :rtype: bool
        """

        slow, fast = n, self.squareOfNumber(n)

        while slow != fast:
            fast = self.squareOfNumber(fast)
            fast = self.squareOfNumber(fast)
            slow = self.squareOfNumber(slow)

        return True if fast == 1 else False

    def squareOfNumber(self, n):
        output = 0

        while n:
            digit = n % 10
            digit = digit ** 2
            output += digit
            n = n // 10
        return output