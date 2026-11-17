# Day 18: Valid Palindrome

## [Question](https://leetcode.com/problems/valid-palindrome/description/)

A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

**Example:**

```
Given s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
```

## Problem-Solving Ideas

## Code
## in Go 

``` Go
func isPalindrome(s string) bool {
    l := 0
    r := len(s)-1
    arr := []rune(s)

    for l < r {
		left := unicode.ToLower(arr[l])
		right := unicode.ToLower(arr[r])

        if !isAlphaNum(left) {
            l ++
            continue
        }

        if !isAlphaNum(right) {
            r --
            continue
        }

        if left != right {
            return false
        }
        l ++
        r --
    }
    return true
}

func isAlphaNum(s rune) bool {
    return unicode.IsLetter(s) || unicode.IsDigit(s)
}
```

## in Python
``` python
class Solution(object):
    def isPalindrome(self, s):
        """
        :type s: str
        :rtype: bool
        """
        l, r = 0, len(s)-1

        while l < r:
            while l < r and not self.isAlphaNum(s[l]):
                l += 1
            while r > l and not self.isAlphaNum(s[r]):
                r -= 1
            if s[l].lower() != s[r].lower():
                return False
            l += 1
            r -= 1
        return True

    def isAlphaNum(self, c):
        return ((ord('A') <= ord(c) <= ord('Z')) or
                (ord('a') <= ord(c) <= ord('z')) or
                (ord('0') <= ord(c) <= ord('9')))
```

## in C++
``` C++
class Solution {
public:
    bool isPalindrome(string s)
    {
        int l = 0;
        int r = s.size()-1;
    
        while(l < r)
        {
            while((l < r) && (!isalnum(s[l])))
            {
                l ++;
            }
            while((l < r) && (!isalnum(s[r])))
            {
                r --;
            }
            if(tolower(s[l]) != tolower(s[r]))
            {
                return false;
            }
            l ++;
            r --;
        }
        return true;
    }
};
```



