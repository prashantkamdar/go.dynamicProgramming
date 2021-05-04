package howSum

func HowSummable(num int, arrNums []int) []int {

	if num == 0 {
		return []int{}
	}

	if num <= 0 {
		return nil
	}

	for _, ele := range arrNums {

		remainder := num - ele
		//fmt.Println(remainder)

		x := HowSummable(remainder, arrNums)
		if x != nil {
			//fmt.Println(x)
			//fmt.Println(reflect.TypeOf(x))
			return x
		}
	}

	return nil

}
