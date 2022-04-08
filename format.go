package utils

import (
	"fmt"
	"strconv"
)

func Float64Format(value float64, num int) float64 {
	format := "%." + strconv.Itoa(num) + "f"
	f, _ := strconv.ParseFloat(fmt.Sprintf(format, value), 64)

	return f
}

func String2Float64(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}
