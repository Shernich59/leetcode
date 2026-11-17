/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
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

	for ptr1 != nil {
		temp.Next = ptr1
		temp = temp.Next
		ptr1 = ptr1.Next
	}

	for ptr2 != nil {
		temp.Next = ptr2
		temp = temp.Next
		ptr2 = ptr2.Next
	}

	dummy = dummy.Next
	return dummy
}