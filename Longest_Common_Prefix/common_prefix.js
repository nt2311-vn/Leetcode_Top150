/**
 * @param {string[]} strs - The array of string to find common
 * @returns {string} returns the common prefix
 */
var longestCommonPrefix = function (strs) {
	const commonPrefix = strs[0];

	for (let i = 0; i < commonPrefix.length; i++) {
		for (let j = 1; j < strs.length; j++) {
			if (i >= strs[j].length || strs[j][i] !== commonPrefix[i]) {
				return commonPrefix.substring(0, i);
			}
		}
	}

	return commonPrefix;
};
