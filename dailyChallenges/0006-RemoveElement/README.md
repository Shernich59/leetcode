# Day 6: Remove element

## [Question](https://leetcode.com/problems/remove-element/description/)

Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The order of the elements may be changed. Then return the number of elements in nums which are not equal to val.

Consider the number of elements in nums which are not equal to val be k, to get accepted, you need to do the following things:

Change the array nums such that the first k elements of nums contain the elements which are not equal to val. The remaining elements of nums are not important as well as the size of nums.
Return k.

**Example:**

```
Given nums = [3,2,2,3], val = 3
Output: 2, nums = [2,2,_,_]
Explanation: Your function should return k = 2, with the first two elements of nums being 2.
It does not matter what you leave beyond the returned k (hence they are underscores).

```

## Problem-Solving Ideas
1. take a k as pointer to monitor on the value in array
2. if the i-th value of nums is not equal to the val, swap itself at k position
3. repeat this with a loop

## Code
## in Go 

``` Go
func removeElement(nums []int, val int) int {
    k := 0

    for i := range(nums){
        if nums[i] != val {
            nums[k] = nums[i]
            k += 1
        }
    }
    return k 
}
```

## in Python

``` python
class Solution(object):
    def removeElement(self, nums, val):
        """
        :type nums: List[int]
        :type val: int
        :rtype: int
        """
        k = 0

        for i in range(len(nums)):
            if nums[i] != val:
                nums[k] = nums[i]
                k += 1
        return k
```