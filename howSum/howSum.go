package howSum

func HowSummable(num int, arrNums []int, x ...*[]int) []int {

	if num == 0 {
		return []int{}
	}

	if num < 0 {
		return nil
	}

	for _, ele := range arrNums {

		remainder := num - ele
		y := HowSummable(remainder, arrNums)
		if y != nil {
			y = append(y, ele)
			return y
		}
	}

	return nil

}
