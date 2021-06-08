package howSum

func HowSummable(num int, arrNums []int, x ...*map[int][]int) []int {

	howSum := make(map[int][]int)

	if len(x) != 0 {
		howSum = *x[0]
	}

	if val, ok := howSum[num]; ok {
		return val
	}

	if num == 0 {
		return []int{}
	}

	if num < 0 {
		return nil
	}

	for _, ele := range arrNums {

		remainder := num - ele
		y := HowSummable(remainder, arrNums, &howSum)
		if y != nil {
			y = append(y, ele)
			howSum[num] = y
			return howSum[num]
		}
	}

	howSum[num] = nil
	return howSum[num]

}
