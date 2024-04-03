package reverselinkedlist

func reverseRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	reverse := reverseRecursion(head.Next)

	head.Next.Next = head
	head.Next = nil

	return reverse
}
