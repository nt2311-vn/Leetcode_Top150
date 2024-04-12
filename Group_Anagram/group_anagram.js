/**
 * @param {string[]} strs - Array of string to find group anagram
 * @returns {string[][]} returns the group anagrams in array
 */
var groupAnagrams = function (strs) {
	const groupAnagrams = new Map();

	for (let i = 0; i < strs.length; i++) {
		let keyStr = strs[i].split("").sort().join("");

		groupAnagrams.has(keyStr)
			? groupAnagrams.get(keyStr).push(strs[i])
			: groupAnagrams.set(keyStr, [strs[i]]);
	}

	const result = [];

	for (const group of groupAnagrams.values()) {
		result.push(group);
	}

	return result;
};
