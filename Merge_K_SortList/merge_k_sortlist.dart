class ListNode {
  int val;
  ListNode? next;
  ListNode([this.val = 0, this.next]);
}

ListNode? mergeSort(ListNode? l1, ListNode? l2) {
  final ListNode dummy = ListNode();
  ListNode? current = dummy;

  while (l1 != null && l2 != null) {
    if (l1.val < l2.val) {
      current!.next = l1;
      l1 = l1.next;
    } else {
      current!.next = l2;
      l2 = l2.next;
    }
    current = current.next!;
  }

  if (l1 != null) {
    current!.next = l1;
  } else {
    current!.next = l2;
  }

  return dummy.next;
}

class Solution {
  ListNode? mergeKLists(List<ListNode?> lists) {
    if (lists.isEmpty) {
      return null;
    }

    if (lists.length == 1) {
      return lists[0];
    }

    int middle = lists.length ~/ 2;

    ListNode? left = mergeKLists(lists.sublist(0, middle));
    ListNode? right = mergeKLists(lists.sublist(middle));

    return mergeSort(left, right);
  }
}
