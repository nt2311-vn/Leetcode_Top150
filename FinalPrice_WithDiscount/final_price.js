/**
 * @param {number[]} prices - The array of price
 * @returns {number[]} return array of final price after discount
 */
var finalPrices = function (prices) {
	let result = new Array(prices.length);

	for (let i = 0; i < prices.length; i++) {
		let discount = 0;

		for (let j = i + 1; j < prices.length; j++) {
			if (prices[j] <= prices[i]) {
				discount = prices[j];
				break;
			}
		}

		result[i] = prices[i] - discount;
	}

	return result;
};
