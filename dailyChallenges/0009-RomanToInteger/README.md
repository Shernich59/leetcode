# Day 9: Roman To Integer

## [Question](https://leetcode.com/problems/roman-to-integer/description/)



**Example:**

```
Given prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
```

## Problem-Solving Ideas
1. using a hashmap
2. check the value before and after
3. if value before is smaller than the next one, substract it
4. else, add it with the next value


## Code
## in Go 

``` Go
func romanToInt(s string) int {
    roman := map[string]int{
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000,
    }

    res := 0

    for i := 0; i < len(s); i++ {
        if i+1 < len(s) && roman[string(s[i])] < roman[string(s[i+1])] {
            res -= roman[string(s[i])]
        } else {
            res += roman[string(s[i])]
        }
    }

    return res
}

```

## in Python
``` python
class Solution(object):
    def romanToInt(self, s):
        """
        :type s: str
        :rtype: int
        """
        roman = {"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

        res = 0

        for i in range(len(s)):
            if i + 1 < len(s) and roman[s[i]] < roman[s[i + 1]]:
                res -= roman[s[i]]
            else:
                res += roman[s[i]]
                
        return res

```