# Day 14: Remove Duplicates From Sorted Array

## [Question](https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/)

Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same. Then return the number of unique elements in nums.

Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:

Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially. The remaining elements of nums are not important as well as the size of nums.
Return k.

**Example:**

```
Given nums = [1,1,2]
Output: 2, nums = [1,2,_]
Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
```

## Problem-Solving Ideas
1. left and right pointer start at the second position in array
Reason: the first element in array must be unique, and no need to replace by the other element
2. compare the element with it's neighbour, if both element are different, update the left pointer with that new element, then increment the left pointer by 1
3. if both are same, increment the right pointer by 1
4. at the end, return the left pointer value, as this will return the number of element in sorted array without duplicates


Time complexity: O(n), as left and right pointer only go thru the entire array once
Space complexity: O(1)

## Code
## in Go 

``` Go
func removeDuplicates(nums []int) int {
    l := 1

    for r := 1; r < len(nums); r++ {
        if nums[r] != nums[r-1] {
            nums[l] = nums[r]
            l ++
        }
    
    }
    return l
}
```

## in Python
``` python
class Solution(object):
    def removeDuplicates(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        l = 1

        for r in range(1, len(nums)):
            if nums[r] != nums[r-1]:
                nums[l] = nums[r]
                l += 1
        return l
         
```

## in C++
``` C++
class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        int l = 1;

        for (int r =1; r < nums.size(); r++)
        {
            if(nums[r] != nums[r-1])
            {
                nums[l] = nums[r];
                l ++;
            }
        }
        return l;
    }
};
```