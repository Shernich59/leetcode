# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

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

        while L1, L2, carry:
            val1 = L1.val if L1 else 0
            val2 = L2.val if L2 else 0

            #New value
            value = val1 + val2 + carry

            #Edge cases: 8 + 7
            carry = value // 10
            value = value % 10
            curr.next = ListNode(value)

            #update pointers
            curr = curr.next
            l1 = l1.next if L1 else None
            L2 = L2.next if L2 else None

        return dummy.next
            


