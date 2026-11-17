class Solution(object):
    def wordPattern(self, pattern, s):
        """
        :type pattern: str
        :type s: str
        :rtype: bool
        """
        word = s.split(" ");
        charToWords, wordsToChar = {}, {}

        if len(pattern) != len(word):
            return False

        for c, w in zip(pattern, word):
            if c in charToWords and charToWords[c] != w:
                return False
            if w in wordsToChar and wordsToChar[w] != c:
                return False
            charToWords[c] = w
            wordsToChar[w] = c

        return True