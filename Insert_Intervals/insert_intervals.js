/**
 * @param {number[][]} intervals - The array of intervals sort in asending order of first element
 * @param {number[]} newInterval - The new interval to insert
 * @returns {number[][]} returns the new intervals after insertion
 */
var insert = function (intervals, newInterval) {
	const result = [];
	let i = 0;

	while (i < intervals.length && intervals[i][1] < newInterval[0]) {
		result.push(intervals[i]);
		i++;
	}

	while (i < intervals.length && intervals[i][0] <= newInterval[1]) {
		newInterval[0] = Math.min(newInterval[0], intervals[i][0]);
		newInterval[1] = Math.max(newInterval[1], intervals[i][1]);
		i++;
	}

	result.push(newInterval);

	while (i < intervals.length) {
		result.push(intervals[i]);
		i++;
	}

	return result;
};
