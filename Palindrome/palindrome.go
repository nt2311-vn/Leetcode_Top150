package palindrome

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	middle, fast := head, head

	for fast != nil {
		middle = middle.Next
		fast = fast.Next

		if fast != nil {
			fast = fast.Next
		}
	}

	var tail *ListNode
	for middle != nil {
		next := middle.Next
		middle.Next = tail
		tail = middle
		middle = next
	}

	for tail != nil {
		if head.Val != tail.Val {
			return false
		}
		head = head.Next
		tail = tail.Next
	}

	return true
}

/*
func nodetoSlice(head *ListNode) []int {
	result := []int{}
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}

	return result
}
*/
