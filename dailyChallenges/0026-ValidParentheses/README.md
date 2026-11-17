# Day 26: Valid Parentheses

## [Question](https://leetcode.com/problems/valid-parentheses/description/?envType=study-plan-v2&envId=top-interview-150)

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.

**Example:**

```
Given s = "()[]{}"
Output: true
```

## Problem-Solving Ideas
1. 
2. 

## Code
## in Go 

``` Go


```

## in Python
``` python
class Solution(object):
    def isValid(self, s):
        """
        :type s: str
        :rtype: bool
        """
        Map = {")": "(", "]": "[", "}": "{"}
        stack = []

        for c in s:
            if c not in Map:
                stack.append(c)
                continue
            if not stack or stack[-1] != Map[c]:
                return False
            stack.pop()

        return not stack
```

## in C++
``` C++


```



