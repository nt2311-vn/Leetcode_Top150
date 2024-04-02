package designdoublylinkedlist

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

type MyLinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func Constructor() MyLinkedList {
	head := &Node{-1, nil, nil}
	tail := &Node{-1, nil, nil}
	head.Next, tail.Prev = tail, head
	return MyLinkedList{
		Head: head,
		Tail: tail,
		Size: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.Size {
		return -1
	}
	cur := this.Head.Next
	for index > 0 {
		cur = cur.Next
		index--
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{
		Val:  val,
		Prev: this.Head,
		Next: this.Head.Next,
	}
	this.Head.Next, this.Head.Next.Prev = node, node
	this.Size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &Node{
		Val:  val,
		Prev: this.Tail.Prev,
		Next: this.Tail,
	}
	this.Tail.Prev, this.Tail.Prev.Next = node, node
	this.Size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Size {
		return
	}
	if index == 0 {
		this.AddAtHead(val)
	} else if index == this.Size {
		this.AddAtTail(val)
	} else {
		cur := this.Head.Next
		for index > 0 {
			cur = cur.Next
			index--
		}
		node := &Node{
			Val:  val,
			Prev: cur.Prev,
			Next: cur,
		}
		cur.Prev.Next, cur.Prev = node, node
		this.Size++
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.Size {
		return
	}
	cur := this.Head.Next
	for index > 0 {
		cur = cur.Next
		index--
	}
	cur.Prev.Next, cur.Next.Prev = cur.Next, cur.Prev
	this.Size--
}
