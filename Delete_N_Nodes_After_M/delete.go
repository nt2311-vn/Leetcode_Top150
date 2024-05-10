package deletennodesafterm

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNodes(head *ListNode, m int, n int) *ListNode {
	sentinel := &ListNode{
		Next: head,
	}
	curr := sentinel
	for curr != nil {
		for i := 0; i < m && curr != nil; i++ {
			curr = curr.Next
		}
		if curr == nil {
			return sentinel.Next
		}

		next := curr
		for i := 0; i < n && next != nil; i++ {
			next = next.Next
		}
		if next == nil {
			curr.Next = nil
			return sentinel.Next
		}
		curr.Next = next.Next
		curr = next
	}
	return sentinel.Next
}
