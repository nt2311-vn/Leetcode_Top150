var MinStack = function () {
	this.stack = [];
	this.minStack = [];
};

/**
 * @param {number} val - Add value to stack
 * @returns {void}
 */
MinStack.prototype.push = function (val) {
	this.stack.push(val);

	if (
		this.minStack.length === 0 ||
		this.minStack[this.minStack.length - 1] >= val
	) {
		this.minStack.push(val);
	}
};

/**
 * @returns {void}
 */
MinStack.prototype.pop = function () {
	if (this.stack.length !== 0) {
		if (
			this.stack[this.stack.length - 1] ===
			this.minStack[this.minStack.length - 1]
		) {
			this.minStack.pop();
		}

		this.stack.pop();
	}
};

/**
 * @returns {number}
 */
MinStack.prototype.top = function () {
	return this.stack[this.stack.length - 1];
};

/**
 * @returns {number}
 */
MinStack.prototype.getMin = function () {
	return this.minStack[this.minStack.length - 1];
};
