# Day 4: Median of Two Sorted Arrays

## [Question](https://leetcode.com/problems/median-of-two-sorted-arrays/description/)


Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).

**Example:**

```
Given nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.

```

## Problem-Solving Ideas
My initial thought:
1. Since both array are sorted, we concat them together (aka merge them)
2. sort the merged array
3. if number of elements in array is even, take the middle and middle + 1 value sum them up, and divide by 2
4. if number of elements in array is odd, (n+1)/2

Time complexity: O(m + n)

TODOS: go with the better solution of O(loh m+n)

## Code
## in Go 

``` Go
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    numTotal := append(nums1, nums2...)

    if len(numTotal) % 2 == 0{
        p1 := (len(numTotal)/2)-1
        p2 := len(numTotal)/2
        return float64(numTotal [p1] + numTotal [p2] )/2.0
    }else{
        return float64 (numTotal [(len(numTotal)+1)/2])
    }
}

```

## in Python

### My initial python code:
``` python
class Solution(object):
    def findMedianSortedArrays(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: float
        """
        numTotal = nums1 + nums2
        numTotal.sort()

        #it is even number
        if len(numTotal) % 2 == 0:
            return (numTotal[(len(numTotal)/2)-1] + numTotal[(len(numTotal)/2)])/2.0

        else:
            pos = (len(numTotal)+1)/2
            return numTotal[pos-1]

```





 

