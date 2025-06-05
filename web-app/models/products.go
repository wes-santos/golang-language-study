package models

import (
	"github.com/wes-santos/alura-golang-study/web-app/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetAllProducts() ([]Product, error) {
	db := db.ConnectToDb()
	defer db.Close()

	query := "SELECT id, name, description, price, quantity FROM products ORDER BY id ASC"
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err = rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Quantity,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func GetProductById(id int) (*Product, error) {
	db := db.ConnectToDb()
	defer db.Close()

	getProductQuery := "SELECT id, name, description, price, quantity FROM products WHERE id = $1"
	row := db.QueryRow(getProductQuery, id)

	var product Product
	err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Quantity,
	)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func AddProduct(product Product) error {
	db := db.ConnectToDb()
	defer db.Close()

	insertStatement, err := db.Prepare(
		"INSERT INTO products (name, description, price, quantity) VALUES ($1, $2, $3, $4)",
	)
	if err != nil {
		return err
	}

	_, err = insertStatement.Exec(product.Name, product.Description, product.Price, product.Quantity)
	if err != nil {
		return err
	}
	defer insertStatement.Close()
	return nil
}

func DeleteProduct(id string) error {
	db := db.ConnectToDb()
	defer db.Close()

	deleteStatement, err := db.Prepare("DELETE FROM products WHERE id = $1")
	if err != nil {
		return err
	}

	_, err = deleteStatement.Exec(id)
	if err != nil {
		return err
	}
	defer deleteStatement.Close()
	return nil
}

func UpdateProduct(product Product) error {
	db := db.ConnectToDb()
	defer db.Close()

	updateStatement, err := db.Prepare(
		"UPDATE products SET name = $1, description = $2, price = $3, quantity = $4 WHERE id = $5",
	)
	if err != nil {
		return err
	}

	_, err = updateStatement.Exec(
		product.Name,
		product.Description,
		product.Price,
		product.Quantity,
		product.Id,
	)
	if err != nil {
		return err
	}
	defer updateStatement.Close()
	return nil
}
