/**
 * @param {number[]} nums
 * @returns {number}
 */
var findTheArrayConcVal = function (nums) {
	let result = 0;
	while (nums.length > 0) {
		let [i, j] = [0, nums.length - 1];

		if (i !== j) {
			let concat = `${nums[i]}${nums[j]}`;
			result += +concat;
		} else {
			result += +nums[i];
		}

		if (nums.length > 1) {
			nums.splice(0, 1);
			nums.splice(nums.length - 1, 1);
		} else {
			break;
		}
	}

	return result;
};
