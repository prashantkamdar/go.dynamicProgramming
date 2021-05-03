package fibonacci

func CalcFib(num int, x ...*map[int]int) int {

	fibSeries := make(map[int]int)

	if len(x) != 0 {
		fibSeries = *x[0]
		//fmt.Println(*x[0])
	}

	if val, ok := fibSeries[num]; ok {
		//fmt.Println(val)
		return val
	}
	if num <= 2 {
		return 1
	}

	fibSeries[num] = CalcFib(num-1, &fibSeries) + CalcFib(num-2, &fibSeries)
	return fibSeries[num]
}
