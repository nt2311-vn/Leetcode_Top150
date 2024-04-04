/**
 * @param {number} n - number of stairs to teach
 * @returns {number} returns the possible way to execute
 */
var climbStairs = function (n) {
	if (n <= 2) {
		return n;
	}

	let first = 1;
	let second = 2;

	for (let i = 3; i <= n; i++) {
		let third = first + second;
		first = second;
		second = third;
	}

	return second;
};
