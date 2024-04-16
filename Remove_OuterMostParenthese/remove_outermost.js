/**
 * @param {string} s - The string of parentheses
 * @returns {string} returns string after remove outer most parentheses
 */
var removeOuterParentheses = function (s) {
	let result = "";
	let pointer = 1;

	for (let i = 1; i < s.length; i++) {
		if (s[i] === "(") {
			pointer++;
		} else {
			pointer--;
		}

		if (pointer === 0) {
			pointer = 1;
			i++;
		} else {
			result += s[i];
		}
	}

	return result;
};
