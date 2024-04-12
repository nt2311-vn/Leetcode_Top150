/**
 * @param {number} n - Factorial of n
 * @returns {number} returs number trail zero of n!
 */
var trailingZeroes = function (n) {
	let count = 0;

	while (n > 0) {
		count += Math.floor(n / 5);
		n /= 5;
	}

	return count;
};
