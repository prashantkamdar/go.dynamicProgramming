package gridTraveller

import (
	"strconv"
)

func CalcPaths(m int, n int, x ...*map[string]int) int {

	paths := make(map[string]int)

	if len(x) != 0 {
		paths = *x[0]
	}

	if val, ok := paths[strconv.Itoa(m)+"+"+strconv.Itoa(n)]; ok {
		return val
	}

	if m == 0 || n == 0 {
		return 0
	}

	if m == 1 && n == 1 {
		return 1
	}

	paths[strconv.Itoa(m)+"+"+strconv.Itoa(n)] = CalcPaths(m-1, n, &paths) + CalcPaths(m, n-1, &paths)
	return paths[strconv.Itoa(m)+"+"+strconv.Itoa(n)]
}
