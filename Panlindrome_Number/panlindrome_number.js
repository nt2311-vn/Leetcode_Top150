/**
 * @param {number} x - Number to checl
 * @returns {boolean} returns isPandom or not
 */
var isPalindrome = function (x) {
	const numStr = x.toString();

	for (let i = 0; i < numStr.length / 2; i++) {
		if (numStr[i] !== numStr[numStr.length - 1 - i]) {
			return false;
		}
	}

	return true;
};
