/**
 * @param {number[]} nums - The array of number
 * @returns {number} returns the number of good pair
 */
var numIdenticalPairs = function (nums) {
	if (nums.length <= 1) {
		return 0;
	}
	let count = 0;
	for (let i = 0; i < nums.length; i++) {
		let j = i + 1;
		while (j < nums.length) {
			if (nums[j] === nums[i]) {
				count++;
			}

			j++;
		}
	}

	return count;
};
