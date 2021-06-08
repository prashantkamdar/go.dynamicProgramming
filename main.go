package main

import (
	"fmt"

	"github.com/prashantkamdar/go.dynamicProgramming/canSum"
	"github.com/prashantkamdar/go.dynamicProgramming/fibonacci"
	"github.com/prashantkamdar/go.dynamicProgramming/gridTraveller"
	"github.com/prashantkamdar/go.dynamicProgramming/howSum"
)

func main() {

	fmt.Println(fibonacci.CalcFib(50))

	fmt.Println(gridTraveller.CalcPaths(18, 18))

	fmt.Println(canSum.Summable(300, []int{7, 14}))

	fmt.Println(howSum.HowSummable(7, []int{2, 3}))
	fmt.Println(howSum.HowSummable(7, []int{5, 3, 4, 7}))
	fmt.Println(howSum.HowSummable(7, []int{2, 4}))
	fmt.Println(howSum.HowSummable(8, []int{2, 3, 5}))
	fmt.Println(howSum.HowSummable(300, []int{7, 14}))
}
