#Day 1: Two Sum

##[Question](https://leetcode.com/problems/two-sum/description/)
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

**Example:**

```
Given nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

```

##Problem-Solving Ideas
Time complexity: O(n)

1. take the difference between target and current value
2. check whether the difference exits in hashmap
3. if value is found, return difference's index and current value index
4. else, store the difference in hashmap and continue to check for the next value

![Overview]()


##Code
##in Go 

``` Go
func twoSum(nums []int, target int) []int {

    m := make(map[int]int)
    for i:=0; i<len(nums); i++{
        diff := target - nums[i]
        if _, ok := m[diff]; ok{
            return []int{m[diff], i}
        }
        m[nums[i]] = i
    }
    return nil
}

```

##in Python

``` python
class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        hashMap = {}

        for i, n in enumerate(nums):
            diff = target - n
            if diff in hashMap:
                return [hashMap[diff], i]
            hashMap[n] = i
            
```