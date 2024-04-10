class Solution {
  List<List<int>> threeSum(List<int> nums) {
    nums.sort((a, b) => a.compareTo(b));
    final List<List<int>> result = [];
    for (var i = 0; i < nums.length; i++) {
      if (i > 0 && nums[i] == nums[i - 1]) {
        continue;
      }

      int j = i + 1;
      int k = nums.length - 1;

      while (j < k) {
        int sum = nums[i] + nums[j] + nums[k];

        if (sum == 0) {
          result.add([nums[i], nums[j], nums[k]]);
          j++;
          k--;
          while (j < k && nums[j - 1] == nums[j]) {
            j++;
          }

          while (j < k && nums[k] == nums[k + 1]) {
            k--;
          }
        } else if (sum > 0) {
          k--;
        } else {
          j++;
        }
      }
    }
    return result;
  }
}
