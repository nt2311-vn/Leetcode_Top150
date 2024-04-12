/**
 * @param {number[]} nums - The array contains integer
 * @returns {number} returns the longest consecutive
 */
var longestConsecutive = function (nums) {
	if (nums.length <= 1) {
		return nums.length;
	}
	const mapNum = new Map();

	for (const num of nums) {
		mapNum.set(num, true);
	}

	let consecutive = 0;

	for (let i = 0; i < nums.length; i++) {
		if (mapNum.has(nums[i] - 1)) {
			continue;
		}

		let sequence = 1;
		let nextSequence = nums[i] + 1;

		while (mapNum.has(nextSequence)) {
			sequence++;
			nextSequence++;
		}

		if (sequence > consecutive) {
			consecutive = sequence;
		}
	}

	return consecutive;
};
