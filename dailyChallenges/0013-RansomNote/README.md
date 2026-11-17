# Day 13: Ransom Note

## [Question](https://leetcode.com/problems/ransom-note/description/)


Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.
Each letter in magazine can only be used once in ransomNote.


**Example:**

```
Given ransomNote = "a", magazine = "b"
Output: false
```

## Problem-Solving Ideas
1. create hashmap for magazine
2. loop thru the ransomNote and check for the count with each character
3. if count of character in magazine is less than that in ransomNote, or 0 return False
3. else decrement the count of that character and continue to check, at the end return True

Time complexity: O(n+m), where n is the length of ransomNote and m is the length of magazine, since we go thru the whole magazine string and ransomNote string.
Space complexity: O(1)

## Code
## in Go 

``` Go
func canConstruct(ransomNote string, magazine string) bool {
    //rune represent int32 type (Unicode character)
    _mag := make(map[rune]int)

    // Populate the _mag with character counts from magazine
    for _, char := range magazine {
        _mag[char] ++
    }

    // Check if the characters in ransomNote can be constructed from magazine
    for _, char := range ransomNote {
        count, found := _mag[char]
        if !found || count == 0 {
            return false
        }
        _mag[char] -- 
    }
    return true
}
```

## in Python
``` python
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
            
        
```

## in C++
``` C++
class Solution {
public:
    bool canConstruct(string ransomNote, string magazine) {
        unordered_map<char, int> _mag;

        // construct a hashmap for magazine
        for(char c : magazine)
        {
            if(_mag.find(c) != _mag.end())
            {
                _mag[c] ++;
            }
            else
            {
                _mag.insert({c,1});
            }
        }
        
        // Iterate for the character of ransom note
        for (char c : ransomNote)
        {
            if(_mag.find(c) == _mag.end()) // element is not found
            {
                return false;
            }
            
            if(_mag[c] == 0)
            {
                return false;
            }
            _mag[c] --;
        }
        return true;
    }
};
```