package canSum

func Summable(num int, arrNums []int, x ...*map[int]bool) bool {

	canSum := make(map[int]bool)

	if len(x) != 0 {
		canSum = *x[0]
	}

	if val, ok := canSum[num]; ok {
		return val
	}

	if num == 0 {
		return true
	}

	if num < 0 {
		return false
	}

	for _, ele := range arrNums {
		remainder := num - ele
		if Summable(remainder, arrNums, &canSum) {
			canSum[num] = true
			return canSum[num]
		}
	}

	canSum[num] = false
	return canSum[num]
}
