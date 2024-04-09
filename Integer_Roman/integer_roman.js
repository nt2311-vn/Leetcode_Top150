/**
 * @param {number} num - Number for converting to Roman
 * @returns {string} returns the Roman string
 */
const numTable = {
	1: "I",
	4: "IV",
	5: "V",
	9: "IX",
	10: "X",
	40: "XL",
	50: "L",
	90: "XC",
	100: "C",
	400: "CD",
	500: "D",
	900: "CM",
	1000: "M",
};
const divideOrder = [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1];
var intToRoman = function (num) {
	let romanStr = "";

	while (num > 0) {
		for (const numInt of divideOrder) {
			let wholeValue = parseInt(num / numInt, 10);
			num -= wholeValue * numInt;
			while (wholeValue > 0) {
				romanStr += numTable[numInt];
				wholeValue--;
			}
		}
	}

	return romanStr;
};
