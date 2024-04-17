/**
 * @param {string} word - The word to reverse
 * @returns {string} returns the word in reversal
 */
const reverse = (word) => {
	let [i, j] = [0, word.length - 1];
	const charArr = word.split("");

	while (i < j) {
		[charArr[i], charArr[j]] = [charArr[j], charArr[i]];
		i++;
		j--;
	}

	return charArr.join("");
};
/**
 * @param {string} s - The string to reverse within word
 * @returns {string} returns the string in original order but within words is reverse
 */
var reverseWords = function (s) {
	const words = s.split(" ");

	for (let i = 0; i < words.length; i++) {
		let reverseWord = reverse(words[i]);
		words[i] = reverseWord;
	}

	return words.join(" ");
};
