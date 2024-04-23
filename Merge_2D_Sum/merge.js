/**
 * @param {number[][]} nums1
 * @param {number[][]} nums2
 * @returns {number[][]}
 */
var mergeArrays = function (nums1, nums2) {
	const result = [];

	let [i, j] = [0, 0];

	while (i < nums1.length && j < nums2.length) {
		if (nums1[i][0] === nums2[j][0]) {
			let sum = nums1[i][1] + nums2[j][1];
			result.push([nums1[i][0], sum]);
			[i, j] = [i + 1, j + 1];
		} else if (nums1[i][0] > nums2[j][0]) {
			result.push(nums2[j]);
			j++;
		} else {
			result.push(nums1[i]);
			i++;
		}
	}

	while (i < nums1.length) {
		result.push(nums1[i]);
		i++;
	}

	while (j < nums2.length) {
		result.push(nums2[j]);
		j++;
	}

	return result;
};
