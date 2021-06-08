package bestSum

func BestSummable(num int, arrNums []int) []int {

	if num == 0 {
		return []int{}
	}

	if num < 0 {
		return nil
	}

	var bestSum []int

	for _, ele := range arrNums {
		remainder := num - ele
		y := BestSummable(remainder, arrNums)

		if y != nil {
			y = append(y, ele)
			if len(bestSum) == 0 || len(y) < len(bestSum) {
				bestSum = y
			}
		}
	}

	return bestSum
}
