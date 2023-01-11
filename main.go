package main

import (
	"fmt"

	"github.com/msharaf315/Learning-GO/models"
)

func main() {
	product := models.NewProduct("1", "FirstProduct", "First Description", 400)
	fmt.Println(product)
}
