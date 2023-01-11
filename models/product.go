package models

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	Id          string
	Title       string
	Description string
	Price       float64
}

func NewProduct(id string, title string, description string, price float64) *Product {
	return &Product{
		id,
		title,
		description,
		price,
	}
}

func (product *Product) PrintData() {
	fmt.Printf(" id: %v,\n title: %v,\n description: %v,\n price: %2f,\n", product.Id, product.Title, product.Description, product.Price)
}

func NewProductFromStdIn() *Product {
	product := Product{}
	fmt.Println("Enter Product Data!")
	fmt.Println("===================")
	reader := bufio.NewReader(os.Stdin)

	product.Id = getUserInput(reader, "Enter Id")

	product.Title = getUserInput(reader, "Enter Title")

	product.Description = getUserInput(reader, "Enter Description")

	product.Price, _ = strconv.ParseFloat(getUserInput(reader, "Enter Price"), 64)

	fmt.Println("Product Created Successfully Product: ")
	product.PrintData()

	return &product
}

func getUserInput(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt + " ")

	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)

	return userInput
}
