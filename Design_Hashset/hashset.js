var MyHashSet = function () {
	this.payload = new Array(10000001).fill(false);
};

/**
 * @param {number} key
 * @returns {void}
 */
MyHashSet.prototype.add = function (key) {
	this.payload[key] = true;
};

/**
 * @param {number} key
 * @returns {void}
 */
MyHashSet.prototype.remove = function (key) {
	this.payload[key] = false;
};

/**
 * @param {number} key
 * @returns {boolean}
 */
MyHashSet.prototype.contains = function (key) {
	return this.payload[key];
};

/**
 * Your MyHashSet object will be instantiated and called as such:
 * var obj = new MyHashSet()
 * obj.add(key)
 * obj.remove(key)
 * var param_3 = obj.contains(key)
 */
