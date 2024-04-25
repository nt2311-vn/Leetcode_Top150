/**
 * @param {number[][]} grid
 * @returns {number}
 */
var countNegatives = function (grid) {
	let result = 0;
	for (const arr of grid) {
		result += binarySearch(arr);
	}

	return result;
};

/**
 * Binanry search to get number of negative integer
 * @param {number[]} target
 * @returns {number}
 */
const binarySearch = (target) => {
	let [i, j] = [0, target.length - 1];

	while (i <= j) {
		let mid = Math.round((i + j) / 2, 0);

		if (target[mid] >= 0) {
			i = mid + 1;
		} else {
			j = mid - 1;
		}
	}

	return target.slice(i).length;
};
