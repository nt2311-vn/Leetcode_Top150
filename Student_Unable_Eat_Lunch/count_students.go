package studentunableeatlunch

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

func (q *Queue) Enqueue(val int) {
	node := &Node{val: val, next: nil}

	if q.IsEmpty() {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}
	q.size++
}

func (q *Queue) Dequeue() *Node {
	if q.IsEmpty() {
		return nil
	}

	dequeueNode := q.head
	q.head = q.head.next

	if q.head == nil {
		q.tail = nil
	}

	q.size--
	return dequeueNode
}

func (q *Queue) Front() int {
	if q.IsEmpty() {
		return -1
	}

	return q.head.val
}

func countStudents(students []int, sandwiches []int) int {
	studentQueue := &Queue{}

	for _, v := range students {
		studentQueue.Enqueue(v)
	}

	for _, v := range sandwiches {
		attemps := studentQueue.size
		for attemps > 0 && studentQueue.Front() != v {
			removed := studentQueue.Dequeue()
			studentQueue.Enqueue(removed.val)
			attemps--
		}

		if attemps == 0 {
			break
		}

		studentQueue.Dequeue()
	}

	return studentQueue.size
}
