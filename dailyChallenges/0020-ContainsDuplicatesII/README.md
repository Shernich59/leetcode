# Day 20: Contains Duplicates II

## [Question](https://leetcode.com/problems/contains-duplicate-ii/)

Given an integer array nums and an integer k, return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.

**Example:**

```
Input: nums = [1,2,3,1], k = 3
Output: true
```

## Problem-Solving Ideas
### Sliding window approach

## Code
## in Go 

``` Go

```

## in Python
``` python
class Solution(object):
    def containsNearbyDuplicate(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: bool
        """
        window = set()
        L = 0

        for R in range(len(nums)):
            if R - L > k:
                window.remove(nums[L])
                L += 1
            if nums[R] in window:
                return True
            window.add(nums[R])
```

## in C++
``` C++

```



