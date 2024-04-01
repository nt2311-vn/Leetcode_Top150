class Solution {
  bool canJump(List<int> nums) {
    var maxSteps = 0;

    for (var i = 0; i < nums.length; i++) {
      if (i > maxSteps) {
        return false;
      }

      if (i + nums[i] > maxSteps) {
        maxSteps = i + nums[i];
      }
    }

    return true;
  }
}
