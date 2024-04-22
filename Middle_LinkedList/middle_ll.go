package middlelinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	oneStep, twoStep := head, head

	for twoStep != nil && twoStep.Next != nil {
		oneStep = oneStep.Next
		twoStep = twoStep.Next.Next
	}

	return oneStep
}
