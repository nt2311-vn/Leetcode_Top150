/**
 * @param {string} word - The word to check and reverse
 * @param {character} ch - The character to index the end of prefix
 * @returns {string} returns the string with reverse prefix if any
 */
var reversePrefix = function (word, ch) {
	let [isFound, index] = [false, 0];
	const charArr = word.split("");

	while (index < charArr.length) {
		if (charArr[index] === ch) {
			isFound = true;
			break;
		}
		index++;
	}

	if (!isFound) {
		return word;
	}

	let [i, j] = [0, index];

	while (i < j) {
		[charArr[i], charArr[j]] = [charArr[j], charArr[i]];
		i++;
		j--;
	}

	return charArr.join("");
};
