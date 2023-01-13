package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	greet()

	// Get starting time
	start := time.Now()
	// Dummy process that takes some time
	storeMoreData(50000, "5000_1.txt")
	storeMoreData(50000, "5000_2.txt")

	// Calculate elapsed time
	elapsedWithoutGoRoutines := time.Since(start)
	fmt.Printf("Without Go Routines took %s \n", elapsedWithoutGoRoutines)

	// Now get calculate time using go routines
	start = time.Now()
	// Dummy process that takes some time

	// initialize the channel
	channel := make(chan int)

	go goStoreMoreData(50000, "5000_1.txt", channel)
	go goStoreMoreData(50000, "5000_2.txt", channel)

	// read the data in the channel
	<-channel
	<-channel
	close(channel)

	// Calculate elapsed time
	elapsedWithGoRoutines := time.Since(start)
	fmt.Printf("With Go Routines took %s\n", elapsedWithGoRoutines)
	timeDelta := (1 - (float64(elapsedWithGoRoutines/time.Millisecond) / float64(elapsedWithoutGoRoutines/time.Millisecond))) * 100

	fmt.Printf("Go Routines are  %.2f %% faster ms Faster\n", timeDelta)
}

func greet() {
	fmt.Println("Hi there!")
}

func storeData(storableText string, fileName string) {
	file, err := os.OpenFile(fileName,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0666,
	)

	if err != nil {
		fmt.Println("Creating the file failed. Exiting.")
		return
	}

	defer file.Close()

	_, err = file.WriteString(storableText)

	if err != nil {
		fmt.Println("Writing to the file failed.")
	}
}

func storeMoreData(lines int, fileName string) {
	for i := 0; i < lines; i++ {
		text := fmt.Sprintf("Line %v - Dummy Data\n", i)
		storeData(text, fileName)
	}

	fmt.Printf("-- Done storing %v lines of text --\n", lines)
}

func goStoreMoreData(lines int, fileName string, c chan int) {
	for i := 0; i < lines; i++ {
		text := fmt.Sprintf("Line %v - Dummy Data\n", i)
		storeData(text, fileName)
	}

	fmt.Printf("-- Done storing %v lines of text --\n", lines)

	//signal to the channel that the function is done by storing a dummy int value to it
	c <- 1

}
