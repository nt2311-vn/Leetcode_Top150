/**
 * @param {number} n - The number to calculate to
 * @returns {number} returns the nth fibonaci number
 */
var fib = function (n) {
	if (n <= 1) {
		return n;
	}

	return fib(n - 1) + fib(n - 2);
};
