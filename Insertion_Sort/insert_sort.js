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
	 * @param {Pair[]} pairs - The tuples of key, value pairs
	 * @returns {Pair[][]} returns the order of each sort process
	 */
	insertionSort(pairs) {
		const results = [];
		for (let i = 0; i < pairs.length; i++) {
			let j = i - 1;
			while (j >= 0 && pairs[j + 1].key < pairs[j].key) {
				[pairs[j], pairs[j + 1]] = [pairs[j + 1], pairs[j]];
				j--;
			}
			results.push([...pairs]);
		}

		return results;
	}
}
