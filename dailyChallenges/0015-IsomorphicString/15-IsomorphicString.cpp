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