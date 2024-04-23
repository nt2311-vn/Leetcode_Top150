/**
 * @param {number[]} nums
 * @returns {number[]}
 */
var sortedSquares = function (nums) {
	const result = [];

	let [i, j] = [0, nums.length - 1];
	while (i <= j) {
		let [iSquare, jSquare] = [nums[i] * nums[i], nums[j] * nums[j]];

		if (iSquare > jSquare) {
			result.push(iSquare);
			i++;
		} else {
			result.push(jSquare);
			j--;
		}
	}

	return result.reverse();
};
