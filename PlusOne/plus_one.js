/**
 * @param {number[]} digits - The array of digits
 * @returns {number[]} returns the array of digits after adding 1 to the number
 */
var plusOne = function (digits) {
	let result = BigInt(digits.join("").toString()) + BigInt("1");
	return result.toString().split("");
};
