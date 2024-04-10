/**
 * @param {number[]} nums - The array of num
 * @param {number} target - The target sum
 * @returns {number[]} returns zero based index two num of which sum = target
 */
var twoSum = function (nums, target) {
	const hashMap = {};

	for (let i = 0; i < nums.length; i++) {
		let remain = target - nums[i];

		if (remain in hashMap) {
			return [hashMap[remain], i];
		} else {
			hashMap[nums[i]] = i;
		}
	}

	return [];
};
