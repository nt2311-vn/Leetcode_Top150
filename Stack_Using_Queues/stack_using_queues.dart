class Node {
  int val;
  Node? next;
  Node(this.val, this.next);
}

class Queue {
  Node? head;
  Node? tail;
  int size = 0;

  Queue(this.head, this.tail, this.size);

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

    Node topItem = head!;
    head = head!.next;

    if (head == null) {
      return tail = null;
    }
    size--;
    return topItem;
  }

  void add(int val) {
    final queueItem = Node(val, null);
    if (isEmpty()) {
      head = queueItem;
      tail = queueItem;
    } else {
      tail!.next = queueItem;
      tail = queueItem;
    }

    size++;
  }
}

class MyStack {
  Queue queue1 = Queue(null, null, 0);
  Queue queue2 = Queue(null, null, 0);

  MyStack();

  void push(int x) {
    queue1.add(x);
  }

  int pop() {
    if (empty()) {
      return -1;
    }
    while (queue1.size > 1) {
      final remove = queue1.remove();
      queue2.add(remove!.val);
    }

    final lastItem = queue1.remove();

    queue2.add(lastItem!.val);

    var temp = queue1;
    queue1 = queue2;
    queue2 = temp;

    return lastItem.val;
  }

  int top() {
    if (empty()) {
      return -1;
    }
    while (queue1.size > 1) {
      final remove = queue1.remove();
      queue2.add(remove!.val);
    }

    final lastItem = queue1.peek();
    queue2.add(queue1.remove()!.val);

    var temp = queue1;
    queue1 = queue2;
    queue2 = temp;

    return lastItem;
  }

  bool empty() {
    return queue1.size == 0 && queue2.size == 0;
  }
}
