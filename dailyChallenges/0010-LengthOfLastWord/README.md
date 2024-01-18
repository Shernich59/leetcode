# Day 10: Length of last word

## [Question](https://leetcode.com/problems/length-of-last-word/description/)


Given a string s consisting of words and spaces, return the length of the last word in the string.

A word is a maximal substring consisting of non-space characters only.

**Example:**

```
Given s = "Hello World"
Output: 5
Explanation: The last word is "World" with length 5.
```

## Problem-Solving Ideas
1. start to search the character from the end of the string
2. increment the length,each time when the string is found
3. return the length of string if a space is found

Time complexity: O(n)
Space complexity: O(1)


## Code
## in Go 

``` Go
func lengthOfLastWord(s string) int {
	l := 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if l >= 1 {
				return l
			}
		} else {
			l++
		}
	}

	return l
}

```

## in Python
``` python
class Solution(object):
    def lengthOfLastWord(self, s):
        """
        :type s: str
        :rtype: int
        """
        i, length = len(s)-1, 0

        while s[i] == " ":
            i -= 1

        while i >= 0 and s[i] != " ":
            length += 1
            i -= 1
        return length

```