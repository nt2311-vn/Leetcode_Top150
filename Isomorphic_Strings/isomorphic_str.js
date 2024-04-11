/**
 * @param {string} s - first string
 * @param {string} t - second string
 * @returns {boolean} returns whether both strings are isomorphic
 */
var isIsomorphic = function (s, t) {
	const firstMap = {};
	const secondMap = {};

	for (let i = 0; i < s.length; i++) {
		if (s[i] in firstMap && firstMap[s[i]] != t[i]) {
			return false;
		}

		firstMap[s[i]] = t[i];

		if (t[i] in secondMap && secondMap[t[i]] != s[i]) {
			return false;
		}

		secondMap[t[i]] = s[i];
	}

	return true;
};
