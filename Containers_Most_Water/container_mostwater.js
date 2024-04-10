/**
 * @param {number[]} height - The height array of water bar
 * @returns {number}  returns the area in integer that contain most water
 */
var maxArea = function (height) {
	let left = 0;
	let right = height.length - 1;
	let area = 0;

	while (left < right) {
		if (height[left] > height[right]) {
			area = Math.max(area, (right - left) * height[right]);
			right--;
		} else {
			area = Math.max(area, (right - left) * height[left]);
			left++;
		}
	}

	return area;
};
