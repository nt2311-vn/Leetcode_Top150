/**
 * @param {string} s - Roman string to convert
 * @returns {number} returns the number in integer
 */
const hashDouble = { IV: 4, IX: 9, XL: 40, XC: 90, CD: 400, CM: 900 };
const hashSingle = { I: 1, V: 5, X: 10, L: 50, C: 100, D: 500, M: 1000 };
var romanToInt = function (s) {
	let result = 0;

	for (let i = 0; i < s.length; i++) {
		if (i + 1 < s.length) {
			if (s.slice(i, i + 2) in hashDouble) {
				result += +hashDouble[s.slice(i, i + 2)];
				i++;
				continue;
			}
		}
		result += +hashSingle[s[i]];
	}
	return result;
};
