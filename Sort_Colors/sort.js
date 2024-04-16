/**
 * @param {number[]} nums - The array of number need to sort
 * @returns {void} Do not return anything, modify nums in-place instead.
 */
var sortColors = function (nums) {
	const bucket = new Array(3).fill(0);

	for (let i = 0; i < nums.length; i++) {
		bucket[nums[i]]++;
	}

	let pointer = 0;
	for (let i = 0; i < bucket.length; i++) {
		while (bucket[i] > 0) {
			nums[pointer] = i;
			pointer++;
			bucket[i]--;
		}
	}
};
