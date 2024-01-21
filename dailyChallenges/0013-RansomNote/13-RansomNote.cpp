class Solution {
public:
    bool canConstruct(string ransomNote, string magazine) {
        unordered_map<char, int> _mag;

        // construct a hashmap for magazine
        for(char c : magazine)
        {
            if(_mag.find(c) != _mag.end())
            {
                _mag[c] ++;
            }
            else
            {
                _mag.insert({c,1});
            }
        }
        
        // Iterate for the character of ransom note
        for (char c : ransomNote)
        {
            if(_mag.find(c) == _mag.end()) // element is not found
            {
                return false;
            }
            
            if(_mag[c] == 0)
            {
                return false;
            }
            _mag[c] --;
        }
        return true;
    }
};