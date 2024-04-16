/**
 * @param {string} s - The string cotains valid parentheses VFS
 * @returns {number} returns the max depth of nest
 */
var maxDepth = function (s) {
	let currentDepth = 0;
	let maxDepth = 0;
	for (let i = 0; i < s.length; i++) {
		if (s[i] === "(") {
			currentDepth++;

			if (currentDepth > maxDepth) {
				maxDepth = currentDepth;
			}
		}

		if (s[i] === ")") {
			currentDepth--;
		}
	}

	return maxDepth;
};
