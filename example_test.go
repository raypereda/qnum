package qnum_test

import (
	"fmt"
	"math"

	"github.com/raypereda/qnum"
)

func Example() {
	n := 1.234567890
	for i := 0; i <= 24; i++ {
		x := n * math.Pow(10.0, float64(i))
		fmt.Printf("%15e %s\n", x, qnum.F(x))
	}
	// Output:
	// 1.234568e+00 1.2
	// 1.234568e+01 12
	// 1.234568e+02 120
	// 1.234568e+03 1.2 Thousand
	// 1.234568e+04 12 Thousand
	// 1.234568e+05 120 Thousand
	// 1.234568e+06 1.2 Million
	// 1.234568e+07 12 Million
	// 1.234568e+08 120 Million
	// 1.234568e+09 1.2 Billion
	// 1.234568e+10 12 Billion
	// 1.234568e+11 120 Billion
	// 1.234568e+12 1.2 Trillion
	// 1.234568e+13 12 Trillion
	// 1.234568e+14 120 Trillion
	// 1.234568e+15 1.2 Quadrillion
	// 1.234568e+16 12 Quadrillion
	// 1.234568e+17 120 Quadrillion
	// 1.234568e+18 1.2 Quintillion
	// 1.234568e+19 12 Quintillion
	// 1.234568e+20 120 Quintillion
	// 1.234568e+21 1.2 Sextillion
	// 1.234568e+22 12 Sextillion
	// 1.234568e+23 120 Sextillion
	// 1.234568e+24 1.2 Septillion
}
