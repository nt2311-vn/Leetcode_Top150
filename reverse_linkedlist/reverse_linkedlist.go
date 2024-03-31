package reverselinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var reverseNode *ListNode

	current := head

	for current != nil {
		remainNode := current.Next
		current.Next = reverseNode
		reverseNode = current
		current = remainNode
	}

	return dummyNode
}
