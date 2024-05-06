class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def getDecimalValue(self, head: ListNode) -> int:
        binary_str = ""
        current = head

        while current is not None:
            binary_str += str(current.val)
            current = current.next

        return int(binary_str, 2)
