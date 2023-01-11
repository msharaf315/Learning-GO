package helper

import (
	"strconv"
	"strings"
)

func ParseStringToFloat(str string) float64{
	str = strings.Replace(str, "\n", "", -1)
	float, _ := strconv.ParseFloat(str, 64)
	
	return float 
}