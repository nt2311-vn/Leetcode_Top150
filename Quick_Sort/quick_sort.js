class Pair {
	/**
	 * @param {number} key - They key to be store in the pair
	 * @param {string} value  - The value to be stored in the pair
	 */
	constructor(key, value) {
		this.key = key;
		this.value = value;
	}
}

class Solution {
	/**
	 * @param {Pair[]} pairs  - The array of Pair to be sorted
	 * @param {number} [start=0] - The beginning index of segment to be sorted
	 * @param {number} [end=pairs.length - 1] - The end index of segment to be sorted
	 * @returns {Pair[]} - The sorted value after quick sort
	 */
	quickSort(pairs, start = 0, end = pairs.length - 1) {
		if (pairs.length <= 1) {
			return pairs;
		}

		if (start < end) {
			let pivot = pairs[end];
			let index = start;

			for (let i = start; i < end; i++) {
				if (pairs[i].key < pivot.key) {
					[pairs[i], pairs[index]] = [pairs[index], pairs[i]];
					index++;
				}
			}

			[pairs[index], pairs[end]] = [pairs[end], pairs[index]];

			this.quickSort(pairs, start, index - 1);
			this.quickSort(pairs, index + 1, end);
		}

		return pairs;
	}
}
