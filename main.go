package main

import (
	"fmt"

	"github.com/msharaf315/Learning-GO/helper"
)



func main(){
	weight, height := getInput();
	
	var weightFloat= helper.ParseStringToFloat(weight)
	var heighttFloat= helper.ParseStringToFloat(height)
	
	bmi := BmiCalculation(weightFloat, heighttFloat)
	fmt.Printf("Your BMI is: %.2f %% \n", bmi)
}