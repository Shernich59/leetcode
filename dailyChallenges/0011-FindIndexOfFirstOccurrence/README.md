# Day 11: Find Index of First Occurrence in a string

## [Question](https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/description)


Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

**Example:**

```
Given haystack = "sadbutsad", needle = "sad"
Output: 0
Explanation: "sad" occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.
```

## Problem-Solving Ideas
### With Naive method (Brute Force solution)
1. Using nested for loop
2. search for the string
3. compare it with needle
4. return 0 if empty
5. return -1 if not found
6. return index if found




Time complexity: O(n*m)
Space complexity: O(1)

### With KNP method

#### /// TODO !!!


## Code
## in Go 

``` Go
func strStr(haystack string, needle string) int {
    if needle == "" {
        return 0
    }

    for i:=0; i< len(haystack) + 1 - len(needle); i++ {
        if haystack[i:i+len(needle)] == needle {
            return i
        }
    }
    return -1
}
```

## in Python
### With Naive method code:
``` python
class Solution(object):
    def strStr(self, haystack, needle):
        """
        :type haystack: str
        :type needle: str
        :rtype: int
        """
        if needle == "":
            return 0
    
        for i in range(len(haystack) + 1 - len(needle)):
            if haystack[i: i + len(needle)] == needle:
                return i
        return -1
    
        
```