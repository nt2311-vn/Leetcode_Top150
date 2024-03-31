package designsinglelinkedlist

type ListNode struct {
	val  int
	next *ListNode
}

type LinkedList struct {
	head   *ListNode
	tail   *ListNode
	length int
}

func constructor() *LinkedList {
	return &LinkedList{head: &ListNode{val: -1}, tail: nil, length: 0}
}

func (this *LinkedList) get(index int) int {
	if index >= this.length {
		return -1
	}

	current := this.head.next

	for i := 0; i < index; i++ {
		if i == index {
			return current.val
		}

		i++

		current = current.next
	}

	return -1
}

func (this *LinkedList) insertHead(val int) {
	newHead := &ListNode{val: val, next: this.head.next}
	this.head.next = newHead
	this.length++

	if this.length == 1 {
		this.tail = newHead
	}
}

func (this *LinkedList) insertTail(val int) {
	newTail := &ListNode{val: val, next: nil}

	if this.tail != nil {
		this.tail.next = newTail
	}
	this.tail = newTail
	if this.length == 0 {
		this.head.next = newTail
	}
	this.length++
}

func (this *LinkedList) remove(index int) bool {
	if index >= this.length || this.length == 0 {
		return false
	}

	var prev *ListNode = this.head

	for i := 0; i < index; i++ {
		prev = prev.next
	}

	if prev.next == this.tail {
		this.tail = prev
	}

	prev.next = prev.next.next

	if prev.next == nil {
		this.tail = prev
	}

	this.length--
	return true
}

func (this *LinkedList) getValues() []int {
	results := []int{}
	current := this.head.next

	for current != nil {
		results = append(results, current.val)
		current = current.next
	}

	return results
}
