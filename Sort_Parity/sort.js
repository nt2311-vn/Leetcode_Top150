/**
 * @param {number[]} nums
 * @return {number[]}
 */
var sortArrayByParity = function (nums) {
	const result = new Array(nums.length);
	let [i, j] = [0, nums.length - 1];

	for (const num of nums) {
		if (num % 2 == 0) {
			result[i] = num;
			i++;
		} else {
			result[j] = num;
			j--;
		}
	}

	return result;
};
