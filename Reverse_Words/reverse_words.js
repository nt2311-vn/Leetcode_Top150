/**
 * @param {string} s - The string to reverse
 * @returns {string} returns words in reversal
 */
var reverseWords = function (s) {
	const words = s.trim().split(/\s+/);

	let i = 0;
	let j = words.length - 1;

	while (i < j) {
		[words[i], words[j]] = [words[j], words[i]];
		i++;
		j--;
	}

	return words.join(" ");
};
