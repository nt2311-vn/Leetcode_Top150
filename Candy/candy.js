/**
 * @param {number[]} ratings - The rate arrays of child
 * @return {number}
 */
var candy = function (ratings) {
	let candies = new Array(ratings.length).fill(1);

	for (let i = 1; i < ratings.length; i++) {
		if (ratings[i - 1] < ratings[i]) {
			candies[i] = candies[i - 1] + 1;
		}
	}

	for (let i = ratings.length - 2; i >= 0; i--) {
		if (ratings[i + 1] < ratings[i] && candies[i + 1] >= candies[i]) {
			candies[i] = candies[i + 1] + 1;
		}
	}

	return candies.reduce((acc, cur) => acc + cur, 0);
};
