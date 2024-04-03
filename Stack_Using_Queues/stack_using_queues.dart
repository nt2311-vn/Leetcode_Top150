class Node {
  int val;
  Node? next;

  Node(this.val, this.next);
}

class Queue {
  Node? head;
  Node? tail;
  int size = 0;

  bool isEmpty() {
    return this.size == 0;
  }

  int peek() {
    if (isEmpty()) {
      return -1;
    }
    return head!.val;
  }

  Node? remove() {
    if (isEmpty()) {
      return null;
    }
    Node? topItem = head;
    head = head?.next;
    if (head == null) {
      tail = null;
    }
    size--;
    return topItem;
  }

  void add(int val) {
    final newNode = Node(val, null);
    if (isEmpty()) {
      head = newNode;
      tail = newNode;
    } else {
      tail?.next = newNode;
      tail = newNode;
    }
    size++;
  }
}

class MyStack {
  Queue queue1 = Queue();
  Queue queue2 = Queue();

  void push(int x) {
    queue1.add(x);
  }

  int pop() {
    if (empty()) {
      return -1;
    }
    while (queue1.size > 1) {
      final Node? removed = queue1.remove();
      queue2.add(removed!.val);
    }
    final Node? lastItem = queue1.remove();
    Queue temp = queue1;
    queue1 = queue2;
    queue2 = temp;
    return lastItem!.val;
  }

  int top() {
    if (empty()) {
      return -1;
    }
    while (queue1.size > 1) {
      queue2.add(queue1.remove()!.val);
    }
    final int lastVal = queue1.peek();
    queue2.add(queue1.remove()!.val);
    Queue temp = queue1;
    queue1 = queue2;
    queue2 = temp;
    return lastVal;
  }

  bool empty() {
    return queue1.isEmpty() && queue2.isEmpty();
  }
}
