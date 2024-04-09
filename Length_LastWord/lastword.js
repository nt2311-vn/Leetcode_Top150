/**
 * @param {string} s - The string input
 * @returns {number} returns the length of lastword
 */
var lengthOfLastWord = function (s) {
	const words = s.trim().split(" ");
	return words[words.length - 1].length;
};
