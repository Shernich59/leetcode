# Day 12: Longest Common Prefix

## [Question](https://leetcode.com/problems/longest-common-prefix/description)


Write a function to find the longest common prefix string amongst an array of strings.
If there is no common prefix, return an empty string "".

**Example:**

```
Given strs = ["flower","flow","flight"]
Output: "fl"
```

## Problem-Solving Ideas
1. using nested for loop
2. search for the first character of each string in the given string list
3. if match, store the character in result and continue step 2
4. if not found, return result

## Code
## in Go 

``` Go
func longestCommonPrefix(strs []string) string {
    var res strings.Builder
    for i := 0; i < len(strs[0]); i++ {
        for _, s := range(strs) {
            if i == len(s) || s[i] != strs[0][i] {
                return res.String()
            }
        }
        res.WriteByte(strs[0][i])
    }
    return res.String()
}

```

## in Python
``` python
class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        res = ""

        for i in range(len(strs)):
            for s in strs:
                if i == len(s) or s[i] != strs[0][i]:
                    return res
            res += strs[0][i]
        return res  
```