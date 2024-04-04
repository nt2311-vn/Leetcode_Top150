/**
 * @param {string} s - The string to check
 * @returns {boolean} Return whether palindrome string is valid
 */
var isPalindrome = function (s) {
	const cleaned = s.toLowerCase().replace(/[^a-z0-9]/g, "");

	let i = 0;
	let j = cleaned.length - 1;

	while (i < j) {
		if (cleaned[i] !== cleaned[j]) {
			return false;
		}

		i++;
		j--;
	}

	return true;
};
