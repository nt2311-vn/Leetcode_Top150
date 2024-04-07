/**
 * Definition for singli-linked list
 * @param {any} val - Value of the list node
 * @param {any} next - Address of next node
 */
function ListNode(val, next) {
	this.val = val === undefined ? 0 : val;
	this.next = next === undefined ? null : next;
}

/**
 * Merge two list nodes into one sorted list node
 * @param {ListNode} l1 - First list node
 * @param {ListNode} l2 - Second list node
 * @returns {ListNode} returns sorted list node
 */
const mergeSort = (l1, l2) => {
	let dummy = new ListNode();
	let current = dummy;

	while (l1 !== null && l2 !== null) {
		if (l1.val < l2.val) {
			current.next = l1;
			l1 = l1.next;
		} else {
			current.next = l2;
			l2 = l2.next;
		}

		current = current.next;
	}

	current.next = l1 || l2;

	return dummy.next;
};

/**
 * @param {ListNode[]} lists - The k lists to sort
 * @returns {ListNode} returns the sorted as singly linked list
 */
var mergeKLists = function (lists) {
	if (lists.length === 0) {
		return null;
	}

	if (lists.length === 1) {
		return lists[0];
	}

	let middle = Math.floor(lists.length / 2);
	let left = mergeKLists(lists.slice(0, middle));
	let right = mergeKLists(lists.slice(middle));

	return mergeSort(left, right);
};
