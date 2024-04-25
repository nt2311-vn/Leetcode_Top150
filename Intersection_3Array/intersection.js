/**
 * @param {number[]} arr1
 * @param {number[]} arr2
 * @param {number[]} arr3
 * @returns {number[]}
 */
var arraysIntersection = function (arr1, arr2, arr3) {
	return arr1
		.filter((x) => binarySearch(x, arr2))
		.filter((x) => binarySearch(x, arr3));
};

/**
 * Binary search an integer in integer array
 * @param {number} val
 * @param {number[]} target
 * @returns {boolean}
 */
const binarySearch = (val, target) => {
	let [i, j] = [0, target.length - 1];

	while (i <= j) {
		let middle = Math.round((i + j) / 2, 0);

		if (target[middle] > val) {
			j = middle - 1;
		} else if (target[middle] < val) {
			i = middle + 1;
		} else {
			return true;
		}
	}

	return false;
};
