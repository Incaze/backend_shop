package store

type Repository struct{}

// Return all products
func (r Repository) GetProducts() Products {
	list := Products{
		Product{
			ID:          0,
			Title:       "Title",
			Image:       "Image",
			Description: "Desc",
			Price:       10000,
			Rating:      5,
		}}
	return list
}

// Return id of product
func (r Repository) GetProductById(id int) Product {
	return Product{}
}

// Add product in the DB
func (r Repository) AddProduct(product Product) bool {
	return false
}

// Update product in the DB
func (r Repository) UpdateProduct(product Product) bool {
	return false
}

// Delete product from DB
func (r Repository) DeleteProduct(id int) string {
	return ""
}
