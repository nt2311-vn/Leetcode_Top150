/**
 * @param {number[][]} image - 2D matrix with bit image data
 * @returns {number[][]} returns the flip horizontal of matrix
 */
var flipAndInvertImage = function (image) {
	if (image.length === 0 || image[0].length === 0) {
		return [];
	}

	const flipImage = [];

	for (let i = 0; i < image.length; i++) {
		flipImage[i] = [];
		for (let j = image[0].length - 1; j >= 0; j--) {
			flipImage[i].push(image[i][j] ^ 1);
		}
	}

	return flipImage;
};
