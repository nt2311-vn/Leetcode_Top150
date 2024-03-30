class MinStack {
  List<int> _stack;
  int _minVal;

  MinStack() {
    _stack = [];
    _minVal = 0;
  }

  void push(int val) {
    _stack.add(val);
  }
}
