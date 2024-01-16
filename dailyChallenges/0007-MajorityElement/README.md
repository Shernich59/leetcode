# Day 7: Majority Element

## [Question](https://leetcode.com/problems/majority-element/description/)


Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

**Example:**

```
Given nums = [3,2,3]
Output: 3
Example 2:

```

**Challenge**
Follow-up ques: Could you solve the problem in linear time and in O(1) space?

## Problem-Solving Ideas
### HashMap approach: 
1. create a hashMap, store the key and value(count) into the hashMap
2. use a variable maxCount to check for the counter
3. return the element if the value of that element exceed maxCount

Time complexity: O(n)
Space complexity: O(n)

### Challenge approach (with Boyer-Moore algorithm):
1. res, count to check on the value and its majority
2. if res is same as the value, increment count by one
3. else, decrement count by one


Time complexity: O(n)
Space complexity: O(1)

## Code
## in Go 
### HashMap approach code:

``` Go
func majorityElement(nums []int) int {
    count := make(map[int]int)
    res, maxCount := 0, 0

    for _,n := range nums{
        count[n] = 1 + count[n]
        if count[n] > maxCount{
            res = n
        }
        maxCount = max(count[n], maxCount)
    }
    return res
}
```
### Challenge approach code (with Boyer-Moore algorithm):
``` Go
func majorityElement(nums []int) int {
    count, res := 0, 0

    for _, n := range(nums){
        if count == 0 {
            res = n
        }
        if res == n {
            count += 1
        }else {
            count -= 1
        }
    }
    return res
}
```

## in Python
### HashMap approach code:
``` python
class Solution(object):
    def majorityElement(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        count = {}
        res, maxCount = 0, 0

        for n in nums:
            count[n] = 1 + count.get(n,0)

            if count[n] > maxCount:
                res = n
            maxCount = max(count[n], maxCount)
        return n
```

### Challenge approach code(with Boyer-Moore algorithm):
```python
class Solution(object):
    def majorityElement(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        count, res = 0, 0

        for n in nums:
            if count == 0:
                res = n
            count += (1 if res == n else -1)
        return ress
```