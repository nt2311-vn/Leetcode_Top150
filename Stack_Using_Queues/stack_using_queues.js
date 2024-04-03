class Item {
	constructor(val) {
		this.val = val;
		this.next = null;
	}
}

class QueueList {
	constructor() {
		this.head = null;
		this.tail = null;
		this.size = 0;
	}

	isEmpty() {
		return this.size === 0;
	}

	add(val) {
		const item = new Item(val);

		if (this.isEmpty()) {
			this.head = item;
			this.tail = item;
		} else {
			this.tail.next = item;
			this.tail = item;
		}

		this.size++;
	}

	remove() {
		if (this.isEmpty()) {
			return null;
		}

		const topItem = this.head;

		this.head = this.head.next;

		if (this.head === null) {
			this.tail = null;
		}
		this.size--;
		return topItem;
	}

	peek() {
		if (this.isEmpty()) {
			return -1;
		}

		return this.head.val;
	}
}

var MyStack = function () {
	this.queue1 = new QueueList();
	this.queue2 = new QueueList();
};

/**
 * @param {number} x
 * @return {void}
 */
MyStack.prototype.push = function (x) {
	this.queue1.add(x);
};

/**
 * @return {number}
 */
MyStack.prototype.pop = function () {
	while (this.queue1.size > 1) {
		const removeItem = this.queue1.remove();
		this.queue2.add(removeItem.val);
	}

	const lastItem = this.queue1.remove();

	this.queue1 = this.queue2;
	this.queue2 = new QueueList();

	return lastItem.val;
};

/**
 * @return {number}
 */
MyStack.prototype.top = function () {
	while (this.queue1.size > 1) {
		const removeItem = this.queue1.remove();
		this.queue2.add(removeItem.val);
	}

	const lastItem = this.queue1.peek();
	this.queue2.add(this.queue1.remove().val);

	this.queue1 = this.queue2;
	this.queue2 = new QueueList();

	return lastItem;
};

/**
 * @return {boolean}
 */
MyStack.prototype.empty = function () {
	return this.queue1.size === 0 && this.queue2.size === 0;
};
