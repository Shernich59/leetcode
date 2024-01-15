# Day 5: Merge sorted array

## [Question](https://leetcode.com/problems/merge-sorted-array/description/)

You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

Merge nums1 and nums2 into a single array sorted in non-decreasing order.

The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

**Example:**

```
Given Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]
Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
Merge result is [1,2,2,3,5,6] with the underlined elements coming from nums1.

```

## Problem-Solving Ideas

### Naive method
1. creating an tmp[] 
2. compare the first value of nums1 and first value of nums2
3. store the smallest value into tmp
4. continue for the next value in both arr

time complexity: O(n)
space complexity: O(n)

### Better approach (to reduce the space complexity into O(1))

```
nums1 [1,2,3,0,0,0]

nums2 [2,5,6]

```
m : no of elements in nums1
n : no of elements in nums2

In order to start from the beginning of nums1, we can start at the back of nums1

```                                                       
nums1 [1,2, ***3*** ,0,0,0]

nums2 [2,5, ***6***]
```

1. compare the largest elements in nums1 and nums2

```
nums1 [1,2, ***3*** ,0,0,***6***]

nums2 [2, ***5***]
```

2. place the largest elements in the last position of nums1,
   repeat steps 1 and steps 2

3. Repeat steps 1 and steps 2 and update the index of nums1 and nums2

```
nums1 [1,2, ***3*** ,0,***5***,6]

nums2 [2]
```
------------------------------------
```
nums1 [1,2, ***3***,5,6]

nums2 [2]
```
-------------------------------------
```
nums1 [1,***2*** ,3,5,6]

nums2 [2]
```
-------------------------------------
```
nums1 [***1*** ,2,3,5,6]

nums2 [***2*** ]
```
-------------------------------------
```
nums1 [1,2,2,3,5,6]

nums2 []
```

3. return nums1

### Edge cases:
if the left nums2 element is greater than first nums1 element

For example:
```
nums1 [***2*** ,2,3,5,6]

nums2 [***1*** ]
```
we need to attach the whole nums2 value to nums1 


time complexity: O(n)
space complexity: O(1)

## Code
## in Go 

``` Go
func merge(nums1 []int, m int, nums2 []int, n int)  {
    //last index in nums1
    last := m + n - 1
    
    for m > 0 && n > 0 {
        if nums2[n-1] > nums1[m-1] {
            nums1[last] = nums2[n-1]
            n--
        }else {
            nums1[last] = nums1[m-1]
            m--
        }
        last--
    }

    for n > 0 {
        nums1[last] = nums2[n-1]
        n--
        last--
    }   
}
```

## in Python

``` python
class Solution(object):
    def merge(self, nums1, m, nums2, n):
        """
        :type nums1: List[int]
        :type m: int
        :type nums2: List[int]
        :type n: int
        :rtype: None Do not return anything, modify nums1 in-place instead.
        """
        #last index in nums1
        last = m + n - 1 

        while m > 0 and n > 0:
            if nums2[n-1] > nums1[m-1]:
                nums1[last] = nums2[n-1]
                n -= 1
            else:
                nums1[last] = nums1[m-1]
                m -= 1
            last -= 1

        #Edge case
        while n > 0:
            nums1[last] = nums2[n-1]
            n -= 1
            last -= 1
```





 

