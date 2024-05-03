/**
 * @param {number} n
 * @returns {number}
 */
var arrangeCoins = function (n) {
	let result = 0;

	for (let i = 1; i <= n; i++) {
		n -= i;
		result++;
		if (n < i) {
			break;
		}
	}

	return result;
};
