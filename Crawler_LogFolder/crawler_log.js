/**
 * @param {string[]} logs - The array of log operations
 * @returns {number} returns the min operation to reach main
 */
var minOperations = function (logs) {
	const stack = [];

	for (let i = 0; i < logs.length; i++) {
		if (logs[i] == "../") {
			if (stack.length > 0) {
				stack.splice(0, 1);
			}
		} else if (logs[i] == "./") {
			continue;
		} else {
			stack.push(null);
		}
	}

	return stack.length;
};
