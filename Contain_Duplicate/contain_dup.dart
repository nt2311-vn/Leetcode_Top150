class Solution {
  bool containsNearbyDuplicate(List<int> nums, int k) {
    final indexTable = {};
    for (var i = 0; i < nums.length; i++) {
      if (indexTable.containsKey(nums[i])) {
        if (i - indexTable[nums[i]] <= k) {
          return true;
        }
      }
      indexTable[nums[i]] = i;
    }

    return false;
  }
}
