package main

import (
	"fmt"

	"github.com/prashantkamdar/go.dynamicProgramming/bestSum"
	"github.com/prashantkamdar/go.dynamicProgramming/canSum"
	"github.com/prashantkamdar/go.dynamicProgramming/fibonacci"
	"github.com/prashantkamdar/go.dynamicProgramming/gridTraveller"
	"github.com/prashantkamdar/go.dynamicProgramming/howSum"
)

func main() {
	fmt.Println("Fibonacci")
	fmt.Println(fibonacci.CalcFib(50))
	fmt.Println("")

	fmt.Println("Grid Traveller")
	fmt.Println(gridTraveller.CalcPaths(18, 18))
	fmt.Println("")

	fmt.Println("Can Sum")
	fmt.Println(canSum.CanSummable(300, []int{7, 14}))
	fmt.Println("")

	fmt.Println("How Sum")
	fmt.Println(howSum.HowSummable(7, []int{2, 3}))
	fmt.Println(howSum.HowSummable(7, []int{5, 3, 4, 7}))
	fmt.Println(howSum.HowSummable(7, []int{2, 4}))
	fmt.Println(howSum.HowSummable(8, []int{2, 3, 5}))
	fmt.Println(howSum.HowSummable(300, []int{7, 14}))
	fmt.Println("")

	fmt.Println("Best Sum")
	fmt.Println(bestSum.BestSummable(7, []int{5, 3, 4, 7}))
	fmt.Println(bestSum.BestSummable(8, []int{2, 3, 5}))
	fmt.Println(bestSum.BestSummable(8, []int{1, 4, 5}))
	fmt.Println(bestSum.BestSummable(100, []int{1, 2, 5, 25}))
	fmt.Println("")
}
