package flippingimage

func flipAndInvertImage(image [][]int) [][]int {
	if len(image) == 0 || len(image[0]) == 0 {
		return [][]int{}
	}
	invertImg := make([][]int, len(image))

	for i := 0; i < len(image); i++ {
		for j := len(image[0]) - 1; j >= 0; j-- {
			invertImg[i] = append(invertImg[i], image[i][j]^1)
		}
	}

	return invertImg
}
