class Solution
{
public:
    bool isHappy(int n) 
    {
        int slow = n;
        int fast = squareOfNum(n);

        while(slow!=fast && fast != 1)
        {
            slow = squareOfNum(slow);
            fast = squareOfNum(squareOfNum(fast));
        }
        if(fast == 1)
        {
            return true;
        }
        return false;
    }
private:
    int squareOfNum(int n)
    {
        int output = 0;
        while(n>0)
        {
            int digit = n % 10;
            digit = pow(digit, 2);
            output += digit;
            n /= 10;
        }
        return output;
    }
};