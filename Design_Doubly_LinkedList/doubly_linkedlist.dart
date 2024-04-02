class MyNode {
  int val;
  MyNode? next;
  MyNode? prev;

  MyNode(this.val, this.next, this.prev);
}

class MyLinkedList {
  MyNode? _head;
  MyNode? _tail;
  int _size = 0;

  MyLinkedList() {
    _head = MyNode(-1, null, null);
    _tail = MyNode(-1, null, null);
    _head!.next = _tail;
    _tail!.prev = _head;
  }
  int get(int index) {
    if (index >= _size) {
      return -1;
    }

    var curr = _head!.next;
    for (var i = 0; i < index; i++) {
      curr = curr!.next;
    }

    return curr!.val;
  }

  void addAtHead(int val) {
    MyNode node = MyNode(val, _head!.next, _head!);

    _head!.next!.prev = node;
    _head!.next = node;
    _size++;
  }

  void addAtTail(int val) {
    MyNode node = MyNode(val, _tail, _tail!.prev);

    _tail!.prev!.next = node;
    _tail!.prev = node;
    _size++;
  }

  void addAtIndex(int index, int val) {
    if (index > _size) {
      return;
    }

    if (index == 0) {
      addAtHead(val);
    } else if (index == _size) {
      addAtTail(val);
    } else {
      MyNode? curr = _head!.next;

      while (index > 0) {
        curr = curr!.next;
        index--;
      }

      MyNode node = MyNode(val, curr!.prev!.next, curr.prev);

      curr.prev!.next = node;
      curr.prev = node;
    }
    _size++;
  }

  void deleteAtIndex(int index) {
    if (index >= _size || index < 0) {
      return;
    }

    MyNode? curr = _head!.next;

    while (index > 0 && curr != null) {
      curr = curr.next;
      index--;
    }
    if (curr != null && curr.next != null && curr.prev != null) {
      curr.prev!.next = curr.next;
      curr.next!.prev = curr.prev;
      _size--;
    }
  }
}
