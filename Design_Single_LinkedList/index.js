class ListNode {
	/**
	 * @param {number} val - Value of a node
	 * @param {ListNode} [nextNode = null] - Reference of the next node
	 */
	constructor(val, nextNode = null) {
		this.val = val;
		this.next = nextNode;
	}
}
class LinkedList {
	constructor() {
		/**
		 * Initialize the list with 'dummy' node, which makes
		 * removing a node from beginning of list easier
		 * @type {ListNode}
		 */
		this.head = new ListNode(-1);
		this.tail = this.head;
		this.length = 0;
	}

	/**
	 * @param {number} index
	 * @returns {number}
	 */
	get(index) {
		if (index >= this.length) {
			return -1;
		}

		let current = this.head.next;
		let i = 0;

		while (current) {
			if (i == index) {
				return current.val;
			}

			i++;
			current = current.next;
		}
	}

	/**
	 * @param {number} val
	 * @returns {void}
	 */
	insertHead(val) {
		let newHead = new ListNode(val);
		let current = this.head.next;
		newHead.next = current;

		this.head.next = newHead;
		this.length++;

		if (!newHead.next) {
			this.tail = newHead;
		}
	}

	/**
	 * @param {number} val
	 * @return {void}
	 */
	insertTail(val) {
		let newTail = new ListNode(val);
		this.tail.next = newTail;
		this.tail = newTail;
		this.length++;
	}

	/**
	 * @param {numer} index
	 * @param {boolean}
	 */
	remove(index) {
		if (index >= this.length || this.length === 0) {
			return false;
		}
		let i = 0;
		let current = this.head;

		while (i < index && current.next) {
			i++;
			current = current.next;
		}

		if (current.next === this.tal) {
			this.tail = current;
		}
		current.next = current.next.next;

		if (current.next === null) {
			this.tail = current;
		}
		this.length--;
		return true;
	}

	/**
	 * @returns {number[]}
	 */
	getValues() {
		let current = this.head.next;
		const results = [];

		while (current) {
			results.push(current.val);
			current = current.next;
		}

		return results;
	}
}

new LinkedList();
