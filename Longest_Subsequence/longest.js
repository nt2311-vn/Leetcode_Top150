/**
 * @param {number[]} nums
 * @param {number[]} queries
 * @returns {number[]}
 */
var answerQueries = function (nums, queries) {
	const result = new Array(queries.length);

	nums.sort((a, b) => a - b);
	for (let j = 0; j < queries.length; j++) {
		let [sum, subsequenceCount] = [0, 0];

		for (let i = 0; i < nums.length; i++) {
			if (sum + nums[i] <= queries[j]) {
				sum += nums[i];
				subsequenceCount++;
			} else {
				break;
			}
		}
		result[j] = subsequenceCount;
	}

	return result;
};
