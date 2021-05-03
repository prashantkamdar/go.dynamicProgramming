package main

import (
	"fmt"

	"github.com/prashantkamdar/go.dynamicProgramming/fibonacci"
	"github.com/prashantkamdar/go.dynamicProgramming/gridTraveller"
)

func main() {

	fmt.Println(fibonacci.CalcFib(50))

	fmt.Println(gridTraveller.CalcPaths(18, 18))
}
