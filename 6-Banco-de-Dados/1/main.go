package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	product := NewProduct("Notebook", 2000.0)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}
	product.Price = 100.0
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	// p, err := selectProduct(db, product.ID)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Product: %v, has price %.2f", p.Name, p.Price)
	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, product := range products {
		fmt.Printf("Product: %v, has price %.2f\n", product.Name, product.Price)
	}

}

func insertProduct(db *sql.DB, product *Product) error {
	steatment, err := db.Prepare("INSERT INTO product(id, name, price) VALUES(?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer steatment.Close()
	_, err = steatment.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	steatment, err := db.Prepare("UPDATE product SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		panic(err)
	}
	defer steatment.Close()
	_, err = steatment.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

// func selectProduct(db *sql.DB, id string) (*Product, error) {
// 	steatment, err := db.Prepare("SELECT id, name, price FROM product WHERE id = ?")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer steatment.Close()
// 	var p Product
// 	err = steatment.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &p, nil
// }

func selectAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
