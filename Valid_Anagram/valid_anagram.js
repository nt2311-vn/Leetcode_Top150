/**
 * @param {string} s - The first string
 * @param {string} t - The second string
 * @returns {boolean} returns anagram relationship
 */
var isAnagram = function (s, t) {
	if (s.length !== t.length) {
		return false;
	}

	const charTable = new Map();

	for (let i = 0; i < s.length; i++) {
		charTable.has(s[i])
			? charTable.set(s[i], charTable.get(s[i]) + 1)
			: charTable.set(s[i], 1);
	}

	for (let i = 0; i < t.length; i++) {
		charTable.has(t[i])
			? charTable.set(t[i], charTable.get(t[i]) - 1)
			: charTable.set(t[i], -1);

		if (charTable.get(t[i]) < 0) {
			return false;
		}
	}

	return true;
};
