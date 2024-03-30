package min_stack

import "math"

type MinStack struct {
	stack      []int
	minElement int
}

func Constructor() MinStack {
	minStack := MinStack{stack: []int{}, minElement: math.MaxInt64}
	return minStack
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)

	if val < this.minElement {
		this.minElement = val
	}
}

func (this *MinStack) Pop() {
	popElement := this.Top()
	this.stack = this.stack[:len(this.stack)-1]

	if this.minElement == popElement {
		this.minElement = math.MaxInt64
		for _, v := range this.stack {
			if v < this.minElement {
				this.minElement = v
			}
		}
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minElement
}
