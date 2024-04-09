import 'dart:math';

class Solution {
  int trap(List<int> height) {
    if (height.length == 0) {
      return 0;
    }
    final bar = height.length;
    final List<int> leftMax = List.filled(bar, 0, growable: false);
    final List<int> rightMax = List.filled(bar, 0, growable: false);

    leftMax[0] = height[0];
    for (var i = 1; i < bar; i++) {
      leftMax[i] = max(leftMax[i - 1], height[i]);
    }

    rightMax[bar - 1] = height[bar - 1];
    for (var i = bar - 2; i >= 0; i--) {
      rightMax[i] = max(rightMax[i + 1], height[i]);
    }

    int water = 0;

    for (var i = 0; i < bar; i++) {
      water += min(leftMax[i], rightMax[i]) - height[i];
    }

    return water;
  }
}
