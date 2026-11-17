# Day 19: Valid Anagram

## [Question](https://leetcode.com/problems/valid-anagram/description/)

Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

**Example:**

```
Given s = "anagram", t = "nagaram"
Output: true
```

## Problem-Solving Ideas
1. create two hashmap T and S
2. compare both length t and s, if different, return false
3. increase the count of character in both S and T
4. compare the character count of both S and T, if different, return false, else return true

## Code
## in Go 

``` Go
func isAnagram(s string, t string) bool {
    sMap := make(map[rune]int)
    tMap := make(map[rune]int)

    if len(s)!= len(t){
        return false
    }

    for _,char := range s{
        sMap[char]++
    }

    for _,char:= range t{
        tMap[char]++
    }

    for key, value1 := range sMap {
        if value2, ok := tMap[key]; !ok || value1 != value2{
            return false
        }
    }
    return true
}
```

## in Python
``` python
class Solution(object):
    def isAnagram(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: bool
        """
        countS , countT = {}, {}

        if (len(s) != len(t)):
            return False

        for i in range(len(s)):
            countS[s[i]] = 1 + countS.get(s[i], 0)
            countT[t[i]] = 1 + countT.get(t[i], 0)

        return countS == countT
```

## in C++
``` C++
class Solution {
public:
    bool isAnagram(string s, string t) {
        unordered_map<char,int> smap;
        unordered_map<char,int> tmap;

        if(s.size() != t.size())
        {
            return false;
        }

        for(int i=0; i < s.size(); i++)
        {
            smap[s[i]]++;
            tmap[t[i]]++;
        }

        for(int i=0; i < smap.size(); i++)
        {
            if(smap[i] != tmap[i])
            {
                return false;
            }
        }
        return true;
    }
};
```



