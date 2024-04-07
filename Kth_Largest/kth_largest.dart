class Solution {
  int findKthLargest(List<int> nums, int k) {
    nums.sort((a, b) => b.compareTo(a));

    return nums[k - 1];
  }
}
