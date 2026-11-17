# Day 3: Longest Substring Without Repeating Characters

## [Question](https://leetcode.com/problems/longest-substring-without-repeating-characters/)

Given a string s, find the length of the longest substring without repeating characters.

**Example:**

```
Given s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

```

## Problem-Solving Ideas

Time complexity: O(n) , go thru the entire string
Space complexity: O(n), because of set(), every value in string maybe unique, so we might need to add all values to the set.
This is similar to 3, 76, 438 and 567, with **'sliding window'concept**.

1. create a set, left pointer and right pointer
2. initialise left pointer to be the first value in string 
3. loop right pointer throughout the string
4. check whether right pointer value exist in set
5. if yes, remove the left pointer value and increment the left pointer
6. else, add the right pointer value into the set
7. increase the length of string


## Code
## in Go 

``` Go
func lengthOfLongestSubstring(s string) int {
    hashSet := make(map[byte]bool)
    l, res := 0, 0

    for r,_ := range s{
        for hashSet[s[r]]{
            delete(hashSet, s[l])
            l ++
        }
        hashSet[s[r]] = true
        res = max(res, r-l+1)
    }
return res

}

func max(a,b int) int {
    if a>b{
        return a
    }
    return b
}

```

## in Python

``` python
class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        hashSet = set()
        l = 0
        res = 0

        for r in range(len(s)):
            while s[r] in hashSet:
                hashSet.remove(s[l])
                l += 1
            hashSet.add(s[r])
            res = max(res, r-l+1)
        return res
```