/**
 * @param {number[]} gas - The gas to fill in each station
 * @param {number[]} cost - cost to move in each station
 * @returns {number} returns the station can complete circuit
 */
var canCompleteCircuit = function (gas, cost) {
	let fill = 0;
	let energy = 0;
	let start = 0;

	for (let i = 0; i < gas.length; i++) {
		fill += gas[i] - cost[i];
		energy += gas[i] - cost[i];

		if (fill < 0) {
			start = i + 1;
			fill = 0;
		}
	}

	if (energy < 0) {
		return -1;
	}

	return start;
};
