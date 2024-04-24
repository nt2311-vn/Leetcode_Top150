/**
 * @param {string} s
 * @param {character} c
 * @returns {number[]}
 */
var shortestToChar = function (s, c) {
	const result = [];

	for (let k = 0; k < s.length; k++) {
		if (s[k] === c) {
			result.push(0);
		} else {
			let [i, j] = [k - 1, k + 1];

			while (i >= 0 && s[i] !== c) {
				i--;
			}

			while (j < s.length && s[j] !== c) {
				j++;
			}

			let minDist = Math.min(k - i, j - k);

			if (i == -1) {
				minDist = j - k;
			}

			if (j == s.length) {
				minDist = k - i;
			}

			result.push(minDist);
		}
	}
	return result;
};
