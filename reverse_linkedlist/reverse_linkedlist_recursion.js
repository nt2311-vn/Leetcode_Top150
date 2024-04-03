/**
 * @param {number} val - The value of node
 * @param {ListNode?} next - The next node address
 */
function ListNode(val, next) {
	this.val = val === undefined ? 0 : val;
	this.next = next === undefined ? null : next;
}

/**
 * @param {ListNode} head - the head list node
 * @returns {ListNode} returns the reverse list node
 */
var reverseList = function (head) {
	if (head === null || head.next === null) {
		return head;
	}

	let reverse = reverseList(head.next);

	head.next.next = head;
	head.next = null;

	return reverse;
};
