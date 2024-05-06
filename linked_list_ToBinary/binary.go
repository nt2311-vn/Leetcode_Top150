package linkedlisttobinary

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	binaryStr := ""

	for head != nil {
		binaryStr += strconv.Itoa(head.Val)
		head = head.Next
	}

	intVal, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		return 0
	}

	return int(intVal)
}
