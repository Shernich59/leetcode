# Day 21: Is Subsequence

## [Question](https://leetcode.com/problems/is-subsequence/description/)

Given two strings s and t, return true if s is a subsequence of t, or false otherwise.

A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).

**Example:**

```
Given s = "abc", t = "ahbgdc"
Output: true
```

## Problem-Solving Ideas
1. using two pointer, i and j to iterate both s and t strings
2. if both s and t value are same, shift the pointer i to the next character in s
3. if both s and t value are not same, shift the pointer j to the next character in t
4. if i reached the end of string s, it means that subsequence is found, return true, else false


## Code
## in Go 

``` Go
func isSubsequence(s string, t string) bool {

    i, j := 0, 0

    if len(s) > len(t) {
        return false
    }

    for i < len(s) {
        for j < len(t) && s[i] != t[j] {
           j ++
        }
        if j == len(t) {
            return false
        }
        i ++ 
        j ++
    }
    return true
}
```

## in Python
``` python
class Solution(object):
    def isSubsequence(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: bool
        """
        i, j = 0 , 0

        if len(s) > len(t):
            return False

        while i < len(s) and j < len(t):
            if s[i] == t[j]:
                i += 1
            j += 1
        return i == len(s) # if i reached the end of s, means it is subsequence of t
```

## in C++
``` C++
class Solution {
public:
    bool isSubsequence(string s, string t) {
        int i = 0;
        int j = 0;

        if(s.length() > t.length())
        {
            return false;
        }

        while ((i < s.length()) && (j < t.length()))
        {
            if(s[i] == t[j])
            {
                i ++;
            }
            j ++;
        }

        return i == s.length();

    }
};
```



