package main

import "fmt"

func main() {
	// 1- init array
	hobbies := [3]string{"coding", "singing", "photography"}

	// 2- print first element alone, second and third element together
	fmt.Printf("first %v, second %v \n", hobbies[0], hobbies[1:3])

	// 3- create two slices in two different ways
	firstSlice := hobbies[:2]

	fmt.Printf("first way: %v,", firstSlice)

	// reslice to contain the second and third element
	firstSlice = firstSlice[1:]

	// create dynamic array
	secondSlice := []string{}
	secondSlice = append(secondSlice, "learn GO")
	secondSlice = append(secondSlice, "Master GO")

	// append with spread operator
	// note we use hobbies[:] to cast it from a list to a slice which can be spread to
	secondSlice = append(secondSlice, hobbies[:]...)

	x := hobbies
	fmt.Println(x)

}
