/**
 * Check if a string is palindrome
 * @param {string} s - String to check
 * @returns {boolean} returns whether s is palindrome
 */
const isPalindrome = (s) => {
	let i = 0;
	let j = s.length - 1;

	while (i < j) {
		if (s[i] !== s[j]) {
			return false;
		}
		i++;
		j--;
	}

	return true;
};
/**
 * @param {string[]} words - The array of word
 * @returns {string} returns the first found palindrome if any
 */
var firstPalindrome = function (words) {
	for (let i = 0; i < words.length; i++) {
		if (isPalindrome(words[i])) {
			return words[i];
		}
	}

	return "";
};
