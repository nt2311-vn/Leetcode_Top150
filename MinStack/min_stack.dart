class MinStack {
  List<int> stack = [];
  List<int> minStack = [];

  MinStack() {}

  void push(int x) {
    stack.add(x);
    if (minStack.isEmpty || x <= minStack.last) {
      minStack.add(x);
    }
  }

  void pop() {
    if (stack.isNotEmpty) {
      if (stack.last == minStack.last) {
        minStack.removeLast();
      }
      stack.removeLast();
    }
  }

  int top() {
    return stack.last;
  }

  int getMin() {
    return minStack.last;
  }
}
