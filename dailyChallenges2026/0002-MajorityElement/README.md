# Day 2: Majority Element

## [Question] (https://leetcode.com/problems/majority-element/description/)

Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

 

**Example 1:**

```
Input: nums = [3,2,3]
Output: 3
```

**Example 2:**

```
Input: nums = [2,2,1,1,1,2,2]
Output: 2
```

```
Constraints:

n == nums.length
1 <= n <= 5 * 104
-109 <= nums[i] <= 109
The input is generated such that a majority element will exist in the array.
 

Follow-up: Could you solve the problem in linear time and in O(1) space?
```
 
## Problem-Solving Ideas

Time complexity: O(n)

1. count the occurences of each value
2. using HashMap method (key:value)
3. key is the number value itself
4. value is the count of each number

## Code
## in Go

``` Go
func majorityElement(nums []int) int {
    count := make(map[int]int)
    result, maxCount := 0, 0

    for _, n := range nums {
        count[n] = count[n] + 1

        if count[n] > maxCount {
            result = n
        }

        if count[n] > maxCount {
            maxCount = count[n]
        }
    }

    return result
}
```

## in Python

``` python
class Solution:
    def majorityElement(self, nums: list[int]) -> int:
        count = {}
        result, maxCount = 0, 0

        for n in nums:
            count[n] = 1 + count.get(n, 0)
            result = n if count[n] > maxCount else result
            maxCount = max(count[n], maxCount)
        return result
```


## Problem-Solving Ideas 2 

Using **Boyer-Moore algorithm**

To be implement