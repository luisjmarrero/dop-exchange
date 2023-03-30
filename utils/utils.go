package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func SetFloatDecimalPoints(x float64, prc int) float64 {
	var factor = math.Pow(10, float64(prc))
	return float64(round(x*factor)) / factor
}

func FormatFloat(num float64, prc int) string {
	var (
		zero, dot = "0", "."

		str = fmt.Sprintf("%."+strconv.Itoa(prc)+"f", num)
	)

	return strings.TrimRight(strings.TrimRight(str, zero), dot)
}
