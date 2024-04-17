/**
 * @param {number[]} nums - The arrays cotain number
 * @param {number} diff - Diff target of each element
 * @returns {number} returns the triplets satisfied substraction of 3 element = diff
 */
var arithmeticTriplets = function (nums, diff) {
	if (nums.length < 3) {
		return 0;
	}
	let triplets = 0;

	for (let i = 0; i < nums.length - 2; i++) {
		let [equalDiff, compareNum] = [0, nums[i]];

		for (let j = i + 1; j < nums.length; j++) {
			if (nums[j] - compareNum == diff) {
				equalDiff++;
				compareNum = nums[j];
			}

			if (equalDiff == 2) {
				triplets++;
				break;
			}
		}
	}

	return triplets;
};
