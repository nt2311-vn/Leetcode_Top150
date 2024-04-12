/**
 * @param {number} target - Target to compare sum
 * @param {number[]} nums - Array of num
 * @returns {number} returns the min length of subarray that sum(subarray) >= target
 */
var minSubArrayLen = function (target, nums) {
	let minLen = nums.length + 1;
	let sum = 0;
	let start = 0;

	for (let end = 0; end < nums.length; end++) {
		sum += nums[end];

		while (sum >= target) {
			if (minLen > end - start + 1) {
				minLen = end - start + 1;
			}
			sum -= nums[start];
			start++;
		}
	}

	if (minLen == nums.length + 1) {
		return 0;
	}

	return minLen;
};
