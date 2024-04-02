class Solution {
  List<int> productExceptSelf(List<int> nums) {
    final _mem = nums.length;
    List<int> _results = List<int>.filled(_mem, 0);

    int leftPointer = 1;

    for (int i = 0; i < _mem; i++) {
      _results[i] = leftPointer;
      leftPointer *= nums[i];
    }

    int rightPointer = 1;

    for (int i = _mem - 1; i >= 0; i--) {
      _results[i] *= rightPointer;
      rightPointer *= nums[i];
    }

    return _results;
  }
}
