from typing import Optional


# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def middleNode(self, head: Optional[ListNode]) -> Optional[ListNode]:
        one_step, two_step = head, head

        while two_step is not None and two_step.next is not None:
            one_step = one_step.next
            two_step = two_step.next.next

        return one_step
