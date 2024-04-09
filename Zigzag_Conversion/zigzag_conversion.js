/**
 * @param {string} s - String to convert
 * @param {number} numRows - number of char in each row
 * @returns {string} returns the string represent in zig zag
 */
var convert = function (s, numRows) {
	if (numRows === 1 || s.length <= numRows) {
		return s;
	}

	let arrChar = new Array(s.length).fill("");
	let zig = true;
	let count = 0;

	for (let i = 0; i < s.length; i++) {
		arrChar[count] += s[i];

		if (zig) {
			count++;
		} else {
			count--;
		}

		if (count === numRows - 1 || count === 0) {
			zig = !zig;
		}
	}

	return arrChar.join("");
};
