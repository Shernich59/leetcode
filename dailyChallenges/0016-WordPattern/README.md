# Day 16: Word Pattern

## [Question](https://leetcode.com/problems/isomorphic-strings/description/)

Given a pattern and a string s, find if s follows the same pattern.

Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s.

**Example:**

```
Given pattern = "abba", s = "dog cat cat dog"
Output: true
```

## Problem-Solving Ideas
1. create two hashmap
2. split the words in string
3. check the length of both pattern and words, if both are different return False
4. map the pattern char to the words
5. if one char is map to many word, or one word is map to different char, return False
6. else, return True

Time complexity: O(n+m), where n is the length of pattern, m is the length of words
Space complexity: O(n+m)

## Code
## in Go 

``` Go
func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	seenChr, seenWord := make(map[byte]int, len(pattern)), make(map[string]int, len(pattern))
	for i := 0; i < len(pattern); i++ {
		chrPos, chrOk := seenChr[pattern[i]]
		wordPos, wordOk := seenWord[words[i]]
		if chrOk != wordOk || chrPos != wordPos {
			return false
		}

		seenChr[pattern[i]] = i
		seenWord[words[i]] = i
	}

	return true
}
```

## in Python
``` python
class Solution(object):
    def wordPattern(self, pattern, s):
        """
        :type pattern: str
        :type s: str
        :rtype: bool
        """
        word = s.split(" ");
        charToWords, wordsToChar = {}, {}

        if len(pattern) != len(word):
            return False

        for c, w in zip(pattern, word):
            if c in charToWords and charToWords[c] != w:
                return False
            if w in wordsToChar and wordsToChar[w] != c:
                return False
            charToWords[c] = w
            wordsToChar[w] = c

        return True
```

## in C++
``` C++
class Solution {
public:
    bool wordPattern(string pattern, string s) {
        std::vector<std::string> words;
        std::istringstream iss(s);
        std::string word;
        std::unordered_map<char, std::string> charToWords;
        std::unordered_map<std::string, char> wordsToChar;

        while(iss >> word)
        {
            words.push_back(word);
        }

        if(pattern.length() != words.size())
        {
            return false;
        }

        for (int i=0; i<pattern.length(); i++)
        {
            char c = pattern[i];
            std::string w = words[i];

            if((charToWords.find(c) != charToWords.end()) && (charToWords[c] != w))
            {
                return false;
            }
            if((wordsToChar.find(w) != wordsToChar.end()) && (wordsToChar[w] != c))
            {
                return false;
            }
            charToWords[c] = w;
            wordsToChar[w] = c;
        }
    return true;
    }
};
```