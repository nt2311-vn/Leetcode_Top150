/**
 * @param {number[][]} intervals - The array of intervals
 * @returns {number[][]} returns the summary merge intervals
 */
var merge = function (intervals) {
	const result = [];

	if (intervals.length === 0) {
		return result;
	}

	intervals.sort((a, b) => a[0] - b[0]);

	result.push(intervals[0]);

	for (let i = 1; i < intervals.length; i++) {
		let current = intervals[i];
		let resultBound = result[result.length - 1];

		if (current[0] <= resultBound[1]) {
			resultBound[1] = Math.max(resultBound[1], current[1]);
		} else {
			result.push(current);
		}
	}

	return result;
};
