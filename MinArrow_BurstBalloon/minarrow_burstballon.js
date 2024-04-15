/**
 * @param {number[][]} points - The array of points position
 * @returns {number} returns the min arrow to burst balloon
 */
var findMinArrowShots = function (points) {
	if (points.length <= 1) {
		return points.length;
	}

	points.sort((a, b) => a[0] - b[0]);

	let count = 1;
	let end = points[0][1];

	for (let i = 0; i < points.length; i++) {
		if (points[i][0] <= end) {
			end = Math.min(points[i][1], end);
		} else {
			count++;
			end = points[i][1];
		}
	}

	return count;
};
