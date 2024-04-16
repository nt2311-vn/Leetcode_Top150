/**
 * @param {string} path - The path need to simplify
 * @returns {string} returns the simplified path
 */
var simplifyPath = function (path) {
	let stack = [];
	let simplifiedPath = "";
	let i = 0;

	while (i < path.length) {
		if (path[i] === "/") {
			i++;
			continue;
		}
		let start = i;
		while (i < path.length && path[i] !== "/") {
			i++;
		}

		let subStr = path.slice(start, i);

		if (subStr == "..") {
			if (stack.length > 0) {
				stack.splice(stack.length - 1, 1);
			}
		} else if (subStr !== ".") {
			stack.push(subStr);
		}
	}

	if (stack.length === 0) {
		return "/";
	}

	for (let j = 0; j < stack.length; j++) {
		simplifiedPath += `/${stack[j]}`;
	}

	return simplifiedPath;
};
