package linkedlistfrequency

type ListNode struct {
	Val  int
	Next *ListNode
}

func frequenciesOfElements(head *ListNode) *ListNode {
	freqMap := map[int]int{}

	for head != nil {
		freqMap[head.Val]++
		head = head.Next
	}

	dummyNode := &ListNode{}
	curr := dummyNode
	for _, val := range freqMap {
		curr.Next = &ListNode{Val: val}
		curr = curr.Next

	}

	return dummyNode.Next
}
