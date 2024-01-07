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

	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
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
        