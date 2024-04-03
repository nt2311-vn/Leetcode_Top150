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
    return size == 0;
  }

  void enQueue(int val) {
    Node node = Node(val, null);
    if (isEmpty()) {
      this.head = node;
      this.tail = node;
    } else {
      this.tail!.next = node;
      this.tail = node;
    }
    this.size++;
  }

  Node? deQueue() {
    if (isEmpty()) {
      return null;
    }

    Node removed = this.head!;
    this.head = this.head!.next;

    if (this.head == null) {
      this.tail == null;
    }
    this.size--;
    return removed;
  }

  int? front() {
    if (isEmpty()) {
      return null;
    }

    return this.head!.val;
  }
}

class Solution {
  int countStudents(List<int> students, List<int> sandwiches) {
    Queue studentQueue = Queue(null, null, 0);

    for (var student in students) {
      studentQueue.enQueue(student);
    }

    for (var sandwich in sandwiches) {
      int maxAttempts = studentQueue.size;

      while (maxAttempts > 0 && studentQueue.front() != sandwich) {
        Node? removed = studentQueue.deQueue();
        studentQueue.enQueue(removed!.val);
        maxAttempts--;
      }

      if (maxAttempts == 0) {
        break;
      }
      studentQueue.deQueue();
    }

    return studentQueue.size;
  }
}
