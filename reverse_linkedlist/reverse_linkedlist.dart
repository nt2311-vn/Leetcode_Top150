class ListNode {
  int val;
  ListNode? next;
  ListNode([this.val = 0, this.next]);
}

class Solution {
  ListNode? reverseList(ListNode? head) {
    ListNode? reverseNode = null;
    var current = head;

    while (current != null) {
      var remainNode = current.next;
      current.next = reverseNode;
      reverseNode = current;
      current = remainNode;
    }

    return reverseNode;
  }
}
