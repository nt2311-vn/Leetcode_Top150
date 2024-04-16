/**
 * @param {number[][]} matrix - The matrix number in non descending order
 * @param {number} target - The target to find
 * @returns {boolean} returns whether target in matrix
 */
var searchMatrix = function (matrix, target) {
	if (matrix.length === 0 || matrix[0].length === 0) {
		return false;
	}

	let [row, col] = [0, matrix[0].length - 1];

	while (row < matrix.length && col >= 0) {
		if (matrix[row][col] > target) {
			col--;
		} else if (matrix[row][col] < target) {
			row++;
		} else {
			return true;
		}
	}

	return false;
};
