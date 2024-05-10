package deletennodesafterm

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNodes(head *ListNode, m int, n int) *ListNode {
	dummyNode := &ListNode{Next: head}
	current := dummyNode

	for current != nil {
		for i := 0; i < m && current != nil && current.Next != nil; i++ {
			current = current.Next
		}

		if current == nil || current.Next == nil {
			break
		}

		next := current.Next

		for i := 0; i < n && next != nil; i++ {
			next = next.Next
		}

		current.Next = next
		current = next
	}
	return dummyNode.Next
}
