# Day 17: Happy Number

## [Question](https://leetcode.com/problems/happy-number/description/)

Write an algorithm to determine if a number n is happy.

A happy number is a number defined by the following process:

Starting with any positive integer, replace the number by the sum of the squares of its digits.
Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
Those numbers for which this process ends in 1 are happy.
Return true if n is a happy number, and false if not.

**Example:**

```
Input: n = 19
Output: true
Explanation:
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1
```

## Problem-Solving Ideas
### Using Floyd's Cycle Detection Algorithm (tortoise and hare algorithm)

#### Normal case:
###### 19 -> 82 -> 68 -> 100 -> 1 -> 1 -> 1 .... 

#### Edge case: 
###### 38 -> **73 -> 58 -> 89 -> 145 -> 51 -> 26 -> 40 -> 16 -> 37 (which for this case back to **) it forms a cycle 

1. create a slow and fast pointer 
2. intialise the slow pointer to the first element, and fast pointer to the element after sum 
3. keep on check for slow and fast pointer, when they meet with each other, means there is a duplicate, and if fast pointer == 1, we return true, else return false

## Code
## in Go 

``` Go
func isHappy(n int) bool {
    slow, fast := n, squareOfNum(n)

    for fast != 1 && slow != fast{
        slow = squareOfNum(slow)
        fast = squareOfNum(squareOfNum(fast))
    }
    return fast == 1
}

func squareOfNum(n int) int {
    output := 0

    for n>0 {
        digit := n % 10
        output += digit * digit
        n /= 10
    }
    return output
}
```

## in Python
``` python
class Solution(object):
    def isHappy(self, n):
        """
        :type n: int
        :rtype: bool
        """

        slow, fast = n, self.squareOfNumber(n)

        while slow != fast:
            fast = self.squareOfNumber(fast)
            fast = self.squareOfNumber(fast)
            slow = self.squareOfNumber(slow)

        return True if fast == 1 else False

    def squareOfNumber(self, n):
        output = 0

        while n:
            digit = n % 10
            digit = digit ** 2
            output += digit
            n = n // 10
        return output
```

## in C++
``` C++
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
```