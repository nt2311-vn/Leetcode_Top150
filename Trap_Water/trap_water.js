/**
 * @param {number[]} height - The array of height bar
 * @returns {number} returns the unit area water can be trapped
 */
var trap = function (height) {
	if (height.length === 0) {
		return 0;
	}
	const bar = height.length;
	const leftMax = new Array(bar);
	const rightMax = new Array(bar);

	leftMax[0] = height[0];

	for (let i = 1; i < bar; i++) {
		leftMax[i] = Math.max(leftMax[i - 1], height[i]);
	}

	rightMax[bar - 1] = height[bar - 1];

	for (let i = bar - 2; i >= 0; i--) {
		rightMax[i] = Math.max(rightMax[i + 1], height[i]);
	}

	let water = 0;

	for (let i = 0; i < bar; i++) {
		water += Math.min(leftMax[i], rightMax[i]) - height[i];
	}

	return water;
};
