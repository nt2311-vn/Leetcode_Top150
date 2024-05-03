/**
 * @param {number} num
 * @returns {boolean}
 */
var isPerfectSquare = function (num) {
	if (num < 4) {
		return num == 1 ? true : false;
	}

	let [i, j] = [0, num / 2];

	while (i <= j) {
		let mid = Math.floor(i + (j - i) / 2);

		if (mid < num / mid) {
			i = mid + 1;
		} else if (mid > num / mid) {
			j = mid - 1;
		} else {
			return true;
		}
	}

	return false;
};
