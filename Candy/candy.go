package candy

func candy(ratings []int) int {
	allocate := make([]int, len(ratings))

	for i := range ratings {
		allocate[i] = 1
	}

	for i := 1; i < len(ratings); i++ {
		if ratings[i-1] < ratings[i] {
			allocate[i] = allocate[i-1] + 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i+1] < ratings[i] && allocate[i+1] >= allocate[i] {
			allocate[i] = allocate[i+1] + 1
		}
	}

	result := 0

	for _, v := range allocate {
		result += v
	}

	return result
}
