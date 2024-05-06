from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def gameResult(self, head: Optional[ListNode]) -> str:
        even_score, odd_score = 0, 0
        while head is not None and head.next is not None:
            even, odd = head.val, head.next.val
            head = head.next.next

            if even > odd:
                even_score += 1
            elif odd > even:
                odd_score += 1

        if even_score > odd_score:
            return "Even"
        elif odd_score > even_score:
            return "Odd"
        else:
            return "Tie"
