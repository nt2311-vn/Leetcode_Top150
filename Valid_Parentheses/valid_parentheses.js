/**
 * @param {string} s - The string to check valid
 * @returns {boolean} returns the boolean value
 */
const hashComplement = { "{": "}", "[": "]", "(": ")" };
var isValid = function (s) {
	if (s.length % 2 !== 0) {
		return false;
	}

	const stack = [];
	for (let i = 0; i < s.length; i++) {
		if (stack.length > 0 && s[i] === hashComplement[stack[stack.length - 1]]) {
			stack.pop();
		} else {
			stack.push(s[i]);
		}
	}

	return stack.length === 0;
};
