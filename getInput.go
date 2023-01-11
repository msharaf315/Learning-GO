package main

import (
	"fmt"

	"github.com/msharaf315/Learning-GO/constants"
)

func getInput()(string, string){
	fmt.Println(constants.Header)
	fmt.Println(constants.Seperator)
	fmt.Print(constants.WeightPrompt)
	weight := readInput()

	fmt.Print(constants.HeightPrompt)
	height := readInput()

	return weight, height;
} 