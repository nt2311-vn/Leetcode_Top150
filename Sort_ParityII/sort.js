/**
 * @param {number[]} nums
 * @returns {number[]}
 */
var sortArrayByParityII = function (nums) {
	const result = new Array(nums.length);

	let [i, j] = [0, 1];

	for (const num of nums) {
		if (num % 2 === 0) {
			result[i] = num;
			i += 2;
		} else {
			result[j] = num;
			j += 2;
		}
	}

	return result;
};
