# Day 15: Isomorphic String

## [Question](https://leetcode.com/problems/isomorphic-strings/description/)

Given two strings s and t, determine if they are isomorphic.

Two strings s and t are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself.

**Example:**

```
Given s = "egg", t = "add"
Output: true
```

## Problem-Solving Ideas
### Using two paper method
1. create two hashmap
2. map the character of s to character of t
3. and vice versa
4. check is the mapping already exist, or is there any two mapping maps to the same character

Time complexity: O(n)

## Code
## in Go 

``` Go
func isIsomorphic(s string, t string) bool {
    mapST := make(map[byte]byte)
    mapTS := make(map[byte]byte)

    if len(s) != len(t) {
		return false
	}

    for i:=0; i < len(s); i++ {
        c1 := s[i]
        c2 := t[i]

        if(mapST[c1] != 0 && mapST[c1] != c2) || (mapTS[c2] != 0 && mapTS[c2] != c1) {
            return false
        }else {
            mapST[c1] = c2
            mapTS[c2] = c1
        }
    }
    return true
}
```

## in Python
``` python
class Solution(object):
    def isIsomorphic(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: bool
        """
        mapST, mapTS = {}, {}

        for c1, c2 in zip(s,t):
            if ((c1 in mapST and mapST[c1] != c2)) or ((c2 in mapTS and mapTS[c2] != c1)):
                return False
            mapST[c1] = c2
            mapTS[c2] = c1
            
        return True

```

## in C++
``` C++
class Solution {
public:
    bool isIsomorphic(string s, string t) {
        unordered_map <char, char> mapST;
        unordered_map <char, char> mapTS;
        char c1, c2;

        if(s.length() != t.length())
        {
            return false;
        }

        for(int i=0; i<s.length(); i++)
        {
            c1 = s[i];
            c2 = t[i];

            if ((mapST.find(c1) != mapST.end() && mapST[c1] != c2) || (mapTS.find(c2) != mapTS.end() && mapTS[c2] != c1)) 
            {
                return false;
            }
            else
            {
                mapST[c1] = c2;
                mapTS[c2] = c1;
            }
        }
        return true;
    }
};
```