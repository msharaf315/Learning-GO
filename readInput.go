package main

import (
	"bufio"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func readInput() string {
	input, _ := reader.ReadString('\n')
	return input
}