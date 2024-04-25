/**
 * @param {number[]} nums
 * @param {number} target
 * @returns {number[]}
 */
var targetIndices = function (nums, target) {
	nums.sort((a, b) => a - b);

	const result = [];

	const { lower, upper } = findBoundary(nums, target);

	for (let i = lower; i <= upper; i++) {
		if (nums[i] === target) {
			result.push(i);
		}
	}

	return result;
};

/**
 * @param {nunber[]} nums
 * @param {number} target
 * @typedef boundaryObj
 * @property {number} lower
 * @property {number} upper
 * @returns {boundaryObj}
 */
const findBoundary = (nums, target) => {
	let [i, j] = [0, nums.length - 1];

	while (i <= j) {
		let mid = Math.round((i + j) / 2, 0);

		if (nums[mid] > target) {
			j = mid - 1;
		} else if (nums[mid] < target) {
			i = mid + 1;
		} else {
			break;
		}
	}

	return { lower: i, upper: j };
};
