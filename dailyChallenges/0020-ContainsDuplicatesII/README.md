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
1. create a hashset, two pointer L and R to create a sliding window
2. add element to hashset, if it is new
3. else, if the sliding window size is larger than k value, remove the element added to hashset, and update the L pointer by 1
4. if element in R pointer found in hashset, return true

## Code
## in Go 

``` Go
func containsNearbyDuplicate(nums []int, k int) bool {
    window := make(map[int]struct{})
    L := 0

    for R:=0; R < len(nums); R++ {
        if R - L > k {
            delete(window, nums[L])
            L ++
        }
        if _, exists := window[nums[R]]; exists {
            return true
        }
        window[nums[R]] = struct{}{}
    }
    return false
}
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
class Solution {
public:
    bool containsNearbyDuplicate(vector<int>& nums, int k) {
        unordered_set<int> window;
        int L = 0;

        for(int R=0; R<nums.size(); R++)
        {
            if((R - L) > k)
            {
                window.erase(nums[L]);
                L ++;
            }
            if(window.find(nums[R]) != window.end())
            {
                return true;
            }
            window.insert(nums[R]);
        }
        return false;
    }
};
```



