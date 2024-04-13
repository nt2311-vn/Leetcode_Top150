/**
 * @param {number} x - Number to find square root
 * @returns {number} returns the square root as loosing integer
 */
var mySqrt = function (x) {
	if (x < 2) {
		return x;
	}

	let left = 0;
	let right = x;

	while (left <= right) {
		let mid = Math.round((left + right) / 2);

		if (mid <= x / mid) {
			left = mid + 1;
		} else {
			right = mid - 1;
		}
	}

	return right;
};
