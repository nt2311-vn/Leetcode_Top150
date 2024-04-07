/**
 * @param {number[]} nums - The array of number to sort
 * @param {number} k - The kth largest
 * @returns {number} returns the required number
 */
var findKthLargest = function (nums, k) {
	nums.sort((a, b) => b - a);
	return nums[k - 1];
};
