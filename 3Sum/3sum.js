/**
 * @param {number[]} nums - The array contain integer
 * @returns {number[][]} returns array of element that have sum = 0
 */
var threeSum = function (nums) {
	const result = [];

	nums.sort((a, b) => a - b);

	for (let i = 0; i < nums.length; i++) {
		if (i > 0 && nums[i] === nums[i - 1]) {
			continue;
		}
		let j = i + 1;
		let k = nums.length - 1;

		while (j < k) {
			let sum = nums[i] + nums[j] + nums[k];

			if (sum === 0) {
				result.push([nums[i], nums[j], nums[k]]);
				j++;
				k--;

				while (j < k && nums[j - 1] === nums[j]) {
					j++;
				}

				while (j < k && nums[k + 1] === nums[k]) {
					k--;
				}
			} else if (sum > 0) {
				k--;
			} else {
				j++;
			}
		}
	}

	return result;
};
