function ListNode(val, next) {
	this.val = val === undefined ? 0 : val;
	this.next = next === undefined ? null : next;
}
/**
 * @param {ListNode} head
 * @returns {ListNode}
 */
var frequenciesOfElements = function (head) {
	const freqMap = new Map();

	while (head !== null) {
		if (!freqMap.has(head.val)) {
			freqMap.set(head.val, 0);
		}
		freqMap.set(head.val, freqMap.get(head.val) + 1);
		head = head.next;
	}

	let dummyNode = new ListNode();
	let curr = dummyNode;

	for (const val of freqMap.values()) {
		curr.next = new ListNode(val);
		curr = curr.next;
	}

	return dummyNode.next;
};
