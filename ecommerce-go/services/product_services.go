package services

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ProductInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var products = []Product{}
var idCounter = 1

func GetAllProducts() []Product {
	return products
}

func CreateProduct(input ProductInput) Product {
	product := Product{
		ID:    idCounter,
		Name:  input.Name,
		Price: input.Price,
	}
	idCounter++
	products = append(products, product)
	return product
}
