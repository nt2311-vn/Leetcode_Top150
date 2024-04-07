class Pair {
	/**
	 * @param {number} key - The key to be stored in the pair
	 * @param {number} value - The value to be stored in the pair
	 */
	constructor(key, value) {
		this.key = key;
		this.value = value;
	}
}

class Solution {
	/**
	 * @param {Pair[]} pairs - The pairs to be sort
	 * @returns {Pair[]} returns sorted pairs
	 */
	mergeSort(pairs) {
		if (pairs.length <= 1) {
			return pairs;
		}

		let middle = Math.floor(pairs.length / 2);

		let left = pairs.slice(0, middle);
		let right = pairs.slice(middle);

		this.mergeSort(left);
		this.mergeSort(right);

		let i = 0,
			j = 0,
			k = 0;

		while (i < left.length && j < right.length) {
			if (left[i].key <= right[j].key) {
				pairs[k++] = left[i++];
			} else {
				pairs[k++] = right[j++];
			}
		}

		while (i < left.length) {
			pairs[k++] = left[i++];
		}

		while (j < right.length) {
			pairs[k++] = right[j++];
		}
		return pairs;
	}
}
