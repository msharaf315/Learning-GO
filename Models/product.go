package models

type Product struct {
	id          string
	title       string
	description string
	price       float64
}

func NewProduct(id string, title string, description string, price float64) *Product {
	return &Product{
		id,
		title,
		description,
		price,
	}

}
