# Day 2: Add Two Numbers

## [Question](https://leetcode.com/problems/add-two-numbers/description/)

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

**Example:**

```
Given l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

```

## Problem-Solving Ideas

**Edge cases:**
**carry forward**

For example
```
Given input: (9 -> 9 -> 9 -> 9 -> 9) + (1 -> )
Output: 0 -> 0 -> 0 -> 0 -> 0 -> 1

```
1. create a dummy head node ( to unify all the test cases)
2. the dummy head node -> next point to the real head


## Code
## in Go 

``` Go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{Val: 0}
    n1, n2, value, carry, curr := 0, 0, 0, 0, dummy

    for l1 != nil || l2 != nil || carry != 0{
        if l1 == nil{
            n1 = 0
        }else{
            n1 = l1.Val
            l1 = l1.Next
        }
        if l2 == nil{
            n2 = 0
        }else{
            n2 = l2.Val
            l2 = l2.Next
        }

        //new value
        value = n1 + n2 + carry
        carry = value / 10
        value = value % 10
        curr.Next = &ListNode{Val: value}

        //update pointers
        curr = curr.Next
    }
    return dummy.Next
}
```

## in Python

``` python
# Definition for singly-linked list.
"""
class ListNode(object):
     def __init__(self, val=0, next=None):
         self.val = val
         self.next = next
"""
class Solution(object):
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """

        dummy = ListNode()
        curr = dummy
        carry = 0

        while l1 or l2 or carry:
            val1 = l1.val if l1 else 0
            val2 = l2.val if l2 else 0

            #New Value
            value = val1 + val2 + carry

            carry = value // 10
            value = value % 10
            curr.next = ListNode(value)

            #Update pointers
            curr = curr.next
            l1 = l1.next if l1 else None
            l2 = l2.next if l2 else None

        return dummy.next
```
