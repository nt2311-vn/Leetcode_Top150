/**
 * @param {string} pattern - The pattern to check
 * @param {string} s - The string words to bijection
 * @returns {boolean} retunrs whether there is bijection relationship
 */
var wordPattern = function (pattern, s) {
	const words = s.split(" ");

	if (words.length !== pattern.length) {
		return false;
	}

	const patternTable = new Map();
	const wordTable = new Map();

	for (let i = 0; i < pattern.length; i++) {
		if (
			patternTable.has(pattern[i]) &&
			patternTable.get(pattern[i]) !== words[i]
		) {
			return false;
		}
		patternTable.set(pattern[i], words[i]);

		if (wordTable.has(words[i]) && wordTable.get(words[i]) !== pattern[i]) {
			return false;
		}

		wordTable.set(words[i], pattern[i]);
	}

	return true;
};
