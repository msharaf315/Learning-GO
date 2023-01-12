package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getUserInput() (string, error) {
	return reader.ReadString('\n')
}

func convertInputToInt(input string) (int, error) {
	input = strings.Replace(input, "\n", "", -1)
	output, err := strconv.Atoi(input)
	return output, err
}

func convertInputToFloat(input string) (float64, error) {
	input = strings.Replace(input, "\n", "", -1)

	output, err := strconv.ParseFloat(input, 0)
	return output, err
}
