/**
 * @param {string} ransomNote - The construct string
 * @param {string} magazine - The whole magazine to check
 * @returns {boolean} returns whether magazine can construct of ransomNote
 */
var canConstruct = function (ransomNote, magazine) {
	const hashMap = {};

	for (let i = 0; i < magazine.length; i++) {
		if (magazine[i] in hashMap) {
			hashMap[magazine[i]]++;
		} else {
			hashMap[magazine[i]] = 1;
		}
	}

	for (let i = 0; i < ransomNote.length; i++) {
		if (ransomNote[i] in hashMap) {
			hashMap[ransomNote[i]]--;
		} else {
			hashMap[ransomNote[i]] = -1;
		}

		if (hashMap[ransomNote[i]] < 0) {
			return false;
		}
	}

	return true;
};
