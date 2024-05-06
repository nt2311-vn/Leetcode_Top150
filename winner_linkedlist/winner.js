function ListNode(val, next) {
	this.val = val === undefined ? 0 : val;
	this.next = next === undefined ? null : next;
}
/**
 * @param {ListNode} head
 * @returns {string}
 */
var gameResult = function (head) {
	let [evenScore, oddScore] = [0, 0];

	while (head !== null && head.next !== null) {
		let [even, odd] = [head.val, head.next.val];
		head = head.next.next;

		if (even > odd) {
			evenScore++;
		} else if (odd > even) {
			oddScore++;
		}
	}

	if (evenScore === oddScore) {
		return "Tie";
	}

	return evenScore > oddScore ? "Even" : "Odd";
};
