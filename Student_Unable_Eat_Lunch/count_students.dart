class Node {
  int val;
  Node? next;
  Node(this.val, this.next);
}

class Queue {
  Node? head;
  Node? tail;
  int size;
  Queue(this.head, this.tail, this.size);

  bool IsEmpty() {
    return size == 0;
  }

  void Enqueue(int val) {
    if (IsEmpty()) {
      Node node = Node(val, null);
      head = node;
      tail = node;
    } else {
      Node node = Node(val, null);
      tail?.next = node;
      tail = node;
    }
    size++;
  }

  Node? Dequeue() {
    if (IsEmpty()) {
      throw Exception("Queue is empty, cannot dequeue");
    }
    Node removed = head!;
    head = head!.next;

    if (head == null) {
      tail = null;
    }
    size--;
    return removed;
  }

  int Front() {
    if (IsEmpty()) {
      throw Exception("Cannot return value of an empty queue");
    }

    return head!.val;
  }
}

class Solution {
  int countStudents(List<int> students, List<int> sandwiches) {
    Queue studentQueue = Queue(null, null, students.length);

    for (var student in students) {
      studentQueue.Enqueue(student);
    }

    for (var sandwich in sandwiches) {
      int maxAttempt = students.length;

      while (maxAttempt > 0 && studentQueue.Front() != sandwich) {
        Node? removed = studentQueue.Dequeue();

        if (removed != null) {
          studentQueue.Enqueue(removed.val);
        }
        maxAttempt--;
      }

      if (maxAttempt == 0) {
        break;
      }

      studentQueue.Dequeue();
    }

    return studentQueue.size;
  }
}
