/**
 * @param {number[]} nums - The array contain integer
 * @returns {number[]} returns the permutation array
 */
var buildArray = function (nums) {
	if (nums.length === 0) {
		return [];
	}

	const result = new Array(nums.length);

	for (let i = 0; i < nums.length; i++) {
		result[i] = nums[nums[i]];
	}

	return result;
};
