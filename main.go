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

	fmt.Println(howSum.HowSummable(7, []int{4, 3}))
}
