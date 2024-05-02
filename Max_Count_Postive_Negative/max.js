/**
 * @param {number[]} nums
 * @returns {number}
 */
var maximumCount = function (nums) {
	const { lower, upper } = binarySearch(nums);
	let [negative, positive] = [0, 0];

	for (let i = 0; i < lower; i++) {
		if (nums[i] < 0) {
			negative++;
		}
	}

	for (let j = upper; j < nums.length; j++) {
		if (j >= 0 && nums[j] > 0) {
			positive++;
		}
	}

	return Math.max(negative, positive);
};

/**
 * Binary search to find lower and upper
 * @param {number[]} numbers - The target number array
 * @typedef {object} boundInfo - Boundary data includes lower and upper bow
 * @property {number} lower - lower boundary
 * @property {number} upper - upper boundary
 * @returns {boundInfo} returns boundary info
 */
const binarySearch = (numbers) => {
	let [i, j] = [0, numbers.length - 1];

	while (i <= j) {
		let mid = Math.floor(i + (j - i) / 2);

		if (numbers[mid] < 0) {
			i = mid + 1;
		} else {
			j = mid - 1;
		}
	}

	return { lower: i, upper: j };
};
