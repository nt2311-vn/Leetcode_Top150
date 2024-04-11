/**
 * @param {number} n - Initial n number
 * @returns {boolean} returns whether n is happy number
 */
var isHappy = function (n) {
	if (n === 1) {
		return true;
	}
	const memoir = new Map();

	while (n !== 1) {
		if (memoir.has(n)) {
			return false;
		}

		memoir.set(n, true);

		let numStr = n.toString();
		let result = 0;

		for (let i = 0; i < numStr.length; i++) {
			result += +numStr[i] * +numStr[i];
		}

		if (result == 1) {
			return true;
		}
		n = result;
	}

	return false;
};
