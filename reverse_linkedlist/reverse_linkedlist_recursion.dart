class ListNode {
  int val;
  ListNode? next;
  ListNode([this.val = 0, this.next]);
}

class Solution {
  ListNode? reverseList(ListNode? head) {
    if (head == null || head.next == null) {
      return head;
    }

    ListNode? reverse = reverseList(head.next);

    head.next!.next = head;
    head.next = null;

    return reverse;
  }
}
