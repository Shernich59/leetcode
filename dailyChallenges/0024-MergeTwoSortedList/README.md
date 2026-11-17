# Day 24: MergeTwoSortedLists

## [Question](https://leetcode.com/problems/merge-two-sorted-lists/description/)

You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.

**Example:**
![Alt text](image.png)
```
Given input list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
```

## Problem-Solving Ideas
1. compare both list value
2. the smaller one assign to the output
3. after comparing the value, in the case where if both list length are not equal, the leftover one will directly link to the end of output
4. dummy node is use as a first head node for the output


A: 1-->2-->3-->4
B: 1-->2--->5-->7-->8

output: 1-->1-->2-->2-->3-->4-->5-->7

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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if(list1 == nil && list2 == nil) {
        return nil
    }
    if(list1 == nil) {
        return list2
    }
    if(list2 == nil) {
        return list1
    }

    ptr1, ptr2 := list1, list2
	dummy := new(ListNode)
	temp := dummy

for ptr1 != nil && ptr2 != nil {
		if ptr1.Val < ptr2.Val {
			temp.Next = ptr1
			temp = temp.Next
			ptr1 = ptr1.Next
		} else {
			temp.Next = ptr2
			temp = temp.Next
			ptr2 = ptr2.Next
		}
	}

	for ptr1 != nil{
		temp.Next = ptr1
		temp = temp.Next
		ptr1 = ptr1.Next
	}

	for ptr2 != nil{
		temp.Next = ptr2
		temp = temp.Next
		ptr2 = ptr2.Next
	}

	dummy = dummy.Next
	return dummy
}
```

## in Python
``` python
# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution(object):
    def mergeTwoLists(self, list1, list2):
        """
        :type list1: Optional[ListNode]
        :type list2: Optional[ListNode]
        :rtype: Optional[ListNode]
        """
        dummy = ListNode()
        tail = dummy

        while list1 and list2:
            if list1.val < list2.val:
                tail.next = list1
                list1 = list1.next
            else:
                tail.next = list2
                list2 = list2.next
            tail = tail.next

        if list1:
            tail.next = list1
        elif list2:
            tail.next = list2
        return dummy.next
```

## in C++
``` C++
/*
    Given heads of 2 sorted linked lists, merge into 1 sorted list
    Ex. list1 = [1,2,4], list2 = [1,3,4] -> [1,1,2,3,4,4]

    Create curr pointer, iterate thru, choose next to be lower one

    Time: O(m + n)
    Space: O(1)
*/

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* mergeTwoLists(ListNode* list1, ListNode* list2) {
        if (list1 == NULL && list2 == NULL) 
        {
            return NULL;
        }
        if (list1 == NULL) 
        {
            return list2;
        }
        if (list2 == NULL) 
        {
            return list1;
        }
        
        ListNode* dummy = new ListNode();
        ListNode *curr = dummy;
        while (list1 != NULL && list2 != NULL) 
        {
            if (list1->val <= list2->val) 
            {
                curr->next = list1;
                list1 = list1->next;
            } else 
            {
                curr->next = list2;
                list2 = list2->next;
            }
            curr = curr->next;
        }
        
        if (list1 == NULL) 
        {
            curr->next = list2;
        } else 
        {
            curr->next = list1;
        }
        
        return dummy->next;
    }
};
```



