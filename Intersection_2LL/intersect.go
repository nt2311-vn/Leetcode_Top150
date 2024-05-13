package intersection2ll

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA := getLength(headA)
	lenB := getLength(headB)

	for lenA > lenB {
		headA = headA.Next
		lenA--
	}

	for lenB > lenA {
		headB = headB.Next
		lenB--
	}

	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}

	return headB
}

func getLength(head *ListNode) int {
	count := 0

	for n := head; n != nil; n = n.Next {
		count++
	}

	return count
}
