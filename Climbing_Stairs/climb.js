/**
 * @param {number} n - number of stairs to teach
 * @returns {number} returns the possible way to execute
 */

const memoir = {};
var climbStairs = function (n) {
	if (n in memoir) {
		return memoir[n];
	}

	if (n <= 2) {
		return n;
	}

	memoir[n] = climbStairs(n - 1) + climbStairs(n - 2);

	return memoir[n];
};
