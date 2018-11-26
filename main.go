package qnum

import (
	"fmt"
	"strconv"
	"strings"
)

var scale map[int]string

// F formats a short scale number with 2 signifigant digits
func F(input float64) string {
	sci := fmt.Sprintf("%3.1e", float64(input)) // example: "1.2e+6"
	splits := strings.Split(sci, "e+")          // example: ["1.2", "6"]
	f, _ := strconv.ParseFloat(splits[0], 64)   // example: 1.2
	exp, _ := strconv.Atoi(splits[1])           // example: 6
	if exp%3 == 1 {
		f *= 10
		exp--
	} else if exp%3 == 2 {
		f *= 100
		exp -= 2
	}

	bignum := scale[exp]
	var digits string
	if f < 10 {
		digits = fmt.Sprintf("%.1f", f)
	} else {
		digits = fmt.Sprintf("%.0f", f)
	}
	return digits + " " + bignum
}

func init() {
	scale = map[int]string{
		2:  "Hundred",
		3:  "Thousand",
		6:  "Million",
		9:  "Billion",
		12: "Trillion",
		15: "Quadrillion",
		18: "Quintillion",
		21: "Sextillion",
		24: "Septillion",
	}
}
