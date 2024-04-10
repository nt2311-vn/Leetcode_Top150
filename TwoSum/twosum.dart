class Solution {
  List<int> twoSum(List<int> nums, int target) {
    final hashMap = {};

    for (var i = 0; i < nums.length; i++) {
      int remain = target - nums[i];
      if (hashMap.containsKey(remain)) {
        return [hashMap[remain], i];
      } else {
        hashMap[nums[i]] = i;
      }
    }

    return [0, 0];
  }
}
