class ListNode {
	/**
	 * @param {number} val - Value of a node
	 * @param {ListNode} [nextNode = null] - Reference of the next node
	 */
	constructor(val, nextNode = null) {
		this.val = val;
		this.nextNode = nextNode;
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
	get(index) {}
}
