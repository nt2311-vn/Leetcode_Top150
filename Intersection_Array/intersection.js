/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @returns {number[]}
 */
var intersection = function (nums1, nums2) {
	nums1.sort((a, b) => a - b);
	nums2.sort((a, b) => a - b);
	const hashMap = new Map();
	const result = [];

	let [i, j] = [0, 0];

	while (i < nums1.length && j < nums2.length) {
		if (nums1[i] === nums2[j]) {
			if (!hashMap.has(nums1[i])) {
				result.push(nums1[i]);
				hashMap.set(nums1[i], true);
			}

			i++;
			j++;
		} else if (nums1[i] > nums2[j]) {
			j++;
		} else {
			i++;
		}
	}

	return result;
};
