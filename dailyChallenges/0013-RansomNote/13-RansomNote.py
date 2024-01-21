class Solution(object):
    def canConstruct(self, ransomNote, magazine):
        """
        :type ransomNote: str
        :type magazine: str
        :rtype: bool
        """
        _mag = Counter(magazine) # <char -> count> hash map

        for c in ransomNote:
            if _mag[c] == 0:
                return False
            else:
                _mag[c] -= 1
        return True
            
        