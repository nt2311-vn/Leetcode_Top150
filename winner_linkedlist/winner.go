package winnerlinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func gameResult(head *ListNode) string {
	oddScore, evenScore := 0, 0

	for head != nil && head.Next != nil {
		even, odd := head.Val, head.Next.Val

		if even > odd {
			evenScore++
		} else if odd > even {
			oddScore++
		}
		head = head.Next.Next
	}

	return findMax(oddScore, evenScore)
}

func findMax(odd, even int) string {
	if odd > even {
		return "Odd"
	} else if even > odd {
		return "Even"
	}

	return "Tie"
}
