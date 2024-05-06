function ListNode(val, next) {
	this.val = val === undefined ? 0 : val;
	this.next = next === undefined ? null : next;
}
/**
 * @param {ListNode} head
 * @returns {number}
 */
var getDecimalValue = function (head) {
	let binaryStr = "";

	while (head !== null) {
		binaryStr += head.val;
		head = head.next;
	}

	return parseInt(binaryStr, 2);
};
