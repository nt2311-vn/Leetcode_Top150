function ListNode(val) {
	this.val = val;
	this.next = null;
}

/**
 * @param {ListNode} headA
 * @param {ListNode} headB
 * @returns {ListNode}
 */
var getIntersectionNode = function (headA, headB) {
	let [a, b] = [headA, headB];

	while (a !== b) {
		a = a === null ? headB : a.next;
		b = b === null ? headA : b.next;
	}

	return a;
};
