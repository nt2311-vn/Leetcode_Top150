/**
 * @param {number[][]} mat
 * @param {number} k
 * @returns {number[]}
 */
var kWeakestRows = function (mat, k) {
	const rowInfo = [];

	for (let i = 0; i < mat.length; i++) {
		rowInfo.push({ rowIndex: i, soldiers: countSoldiers(mat[i]) });
	}

	rowInfo.sort((a, b) => {
		if (a.soldiers === b.soldiers) {
			return a.rowIndex - b.rowIndex;
		}
		return a.soldiers - b.soldiers;
	});

	const result = [];

	for (let i = 0; i < k; i++) {
		result.push(rowInfo[i].rowIndex);
	}

	return result;
};

/**
 * @param {number[]} row
 * @returns {number}
 */
const countSoldiers = (row) => {
	let [i, j] = [0, row.length - 1];

	while (i <= j) {
		let mid = Math.floor((i + j) / 2);
		if (row[mid] == 1) {
			i = mid + 1;
		} else {
			j = mid - 1;
		}
	}

	return i;
};
