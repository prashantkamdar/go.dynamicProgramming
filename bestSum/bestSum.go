package bestSum

func BestSummable(num int, arrNums []int, x ...*map[int][]int) []int {

	memoBestSum := make(map[int][]int)

	if len(x) != 0 {
		memoBestSum = *x[0]
	}

	if val, ok := memoBestSum[num]; ok {
		return val
	}

	if num == 0 {
		return []int{}
	}

	if num < 0 {
		return nil
	}

	var bestSum []int

	for _, ele := range arrNums {
		remainder := num - ele
		y := BestSummable(remainder, arrNums, &memoBestSum)

		if y != nil {
			y = append(y, ele)
			if len(bestSum) == 0 || len(y) < len(bestSum) {
				bestSum = y
			}
		}
	}

	memoBestSum[num] = bestSum
	return memoBestSum[num]
}
