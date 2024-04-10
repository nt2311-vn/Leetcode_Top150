import 'dart:math';

class Solution {
  int maxArea(List<int> height) {
    int left = 0;
    int right = height.length - 1;
    int area = 0;

    while (left < right) {
      if (height[left] > height[right]) {
        area = max(area, (right - left) * height[right]);
        right--;
      } else {
        area = max(area, (right - left) * height[left]);
        left++;
      }
    }

    return area;
  }
}
