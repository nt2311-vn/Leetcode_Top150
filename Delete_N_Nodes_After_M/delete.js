function ListNode(val, next) {
	this.val = val === undefined ? 0 : val;
	this.next = next === undefined ? null : next;
}
/**
 * @param {ListNode} head
 * @param {number} m
 * @param {number} n
 * @returns {ListNode}
 */
var deleteNodes = function (head, m, n) {
	const dummyNode = new ListNode(undefined, head);
	let currentNode = dummyNode;

	while (currentNode !== null) {
		for (let i = 0; i < m && currentNode !== null; i++) {
			currentNode = currentNode.next;
		}

		if (currentNode === null) {
			return dummyNode.next;
		}

		let nextNode = currentNode;

		for (let i = 0; i < n && nextNode !== null; i++) {
			nextNode = nextNode.next;
		}

		if (nextNode === null) {
			currentNode.next = null;
			return dummyNode.next;
		}

		currentNode.next = nextNode.next;
		currentNode = nextNode;
	}

	return dummyNode.next;
};
