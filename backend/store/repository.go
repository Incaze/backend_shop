package store

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Repository struct{}

const (
	DBNAME   = "shopDB"
	USERNAME = "root"
	PASSWORD = "hesoyam"
	DRIVER   = "mysql"
	SERVER   = USERNAME + ":" + PASSWORD + "@/" + DBNAME
)

// Return all products
func (r Repository) GetProducts() Products {
	db, err := sql.Open(DRIVER, SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to MySQL server:", err)
	}
	defer db.Close()

	rows, err := db.Query("select * from `products`")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	products := Products{}
	for rows.Next() {
		p := Product{}
		err := rows.Scan(&p.ID, &p.Description, &p.Image, &p.Price, &p.Rating, &p.Title)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}
	return products
}

// Return product by id
func (r Repository) GetProductById(id int) Product {
	db, err := sql.Open(DRIVER, SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to MySQL server:", err)
	}
	defer db.Close()

	rows, err := db.Query("select * from `products` where `id` = ?", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	product := Product{}
	for rows.Next() {
		err = rows.Scan(&product.ID, &product.Description, &product.Image, &product.Price, &product.Rating, &product.Title)
		if err != nil {
			panic(err)
		}
	}
	return product
}

// Add product in the table
func (r Repository) AddProduct(product Product) bool {
	db, err := sql.Open(DRIVER, SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to MySQL server:", err)
	}
	defer db.Close()

	_, err = db.Query("insert into `products` "+
		"(title, rating, price, image, description) values (?, ?, ?, ?, ?)",
		product.Title, product.Rating, product.Price, product.Image, product.Description)

	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// Update product in the DB
func (r Repository) UpdateProduct(product Product) bool {
	db, err := sql.Open(DRIVER, SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to MySQL server:", err)
	}
	defer db.Close()

	_, err = db.Query("update `products` set "+
		"title = ?, rating = ?, price = ?, image = ?, description = ? where id = ?",
		product.Title, product.Rating, product.Price, product.Image, product.Description, product.ID)

	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// Delete product from DB
func (r Repository) DeleteProduct(id int) string {
	db, err := sql.Open(DRIVER, SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to MySQL server:", err)
	}
	defer db.Close()

	_, err = db.Query("delete from `products` where id = ?", id)

	if err != nil {
		fmt.Println(err)
		return "INTERNAL ERR"
	}
	return "OK"
}
