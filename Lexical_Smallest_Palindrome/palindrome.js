/**
 * @param {string} s - The string to make palindrome
 * @returns {string} returns the palindrome with smallest lexical modification
 */
var makeSmallestPalindrome = function (s) {
	const charArr = s.split("");

	let [i, j] = [0, charArr.length - 1];

	while (i < j) {
		if (charArr[i] !== charArr[j]) {
			charArr[i] = charArr[i] > charArr[j] ? charArr[j] : charArr[i];
			charArr[j] = charArr[i] > charArr[j] ? charArr[j] : charArr[i];
		}
		i++;
		j--;
	}

	return charArr.join("");
};
