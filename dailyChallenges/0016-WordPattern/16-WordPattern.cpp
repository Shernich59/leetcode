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