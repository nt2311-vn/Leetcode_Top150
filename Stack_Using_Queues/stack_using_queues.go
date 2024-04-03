package stackusingqueues

type Node struct {
	val  int
	next *Node
}

type Queue struct {
	head *Node
	tail *Node
	size int
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Add(val int) {
	node := &Node{val: val}
	if q.IsEmpty() {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}

	q.size++
}

func (q *Queue) Peek() int {
	if q.IsEmpty() {
		return -1
	}

	return q.head.val
}

func (q *Queue) Pop() *Node {
	if q.IsEmpty() {
		return nil
	}

	removed := q.head
	q.head = q.head.next

	if q.head == nil {
		q.tail = nil
	}

	q.size--
	return removed
}

type MyStack struct {
	queue1 *Queue
	queue2 *Queue
}

func Constructor() MyStack {
	return MyStack{queue1: &Queue{}, queue2: &Queue{}}
}

func (this *MyStack) Push(x int) {
}

func (this *MyStack) Pop() int {
}

func (this *MyStack) Top() int {
}

func (this *MyStack) Empty() bool {
	return this.queue1.size == 0 && this.queue2.size == 0
}
