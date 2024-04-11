/**
 * @param {number[]} nums - The array cotain number
 * @param {number} k - k position to check
 * @returns {boolean} returns the satisfied check
 */
var containsNearbyDuplicate = function (nums, k) {
	const indexTable = new Map();

	for (let i = 0; i < nums.length; i++) {
		if (indexTable.has(nums[i])) {
			const existIndex = indexTable.get(nums[i]);
			if (i - existIndex <= k) {
				return true;
			}
		}

		indexTable.set(nums[i], i);
	}

	return false;
};
