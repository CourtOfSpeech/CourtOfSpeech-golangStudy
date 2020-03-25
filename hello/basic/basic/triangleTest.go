package main

import (
	"fmt"
)

type trangle struct {
	a, b, c int
}

func testTringle() (a int, err error) {
	trangles := []trangle{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 333},
		{30000, 40000, 50000},
	}

	for _, trangle := range trangles {
		if actual := calcTriangle(trangle.a, trangle.b); actual != trangle.c {
			err = fmt.Errorf("calcTriangle(%d, %d);"+
				"got %d; expected %d", trangle.a, trangle.b, actual, trangle.c)
			return
		}
	}
	return
}
