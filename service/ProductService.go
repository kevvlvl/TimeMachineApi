package service

import (
	"TimeMachineApi/database"
	"fmt"
)

type Product struct {
	id           int    `json:"id"`
	productName  string `json:"productName"`
	productType  string `json:"productType"`
	supplierCost int    `json:"supplierCost"`
	salePrice    int    `json:"salePrice"`
}

func GetProducts() []Product {

	fmt.Println("BEGIN GetProducts")

	db := database.OpenConnection()
	rows, _ := db.Query("SELECT id, productName, productType, supplierCost, salePrice FROM product")
	var products []Product

	if rows != nil {

		for rows.Next() {
			var id int
			var productName string
			var productType string
			var supplierCost int
			var retailPrice int

			err := rows.Scan(&id, &productName, &productType, &supplierCost, &retailPrice)

			if err != nil {
				fmt.Println("Error trying to scan row for the listed fields: ", err)
				return nil
			}

			products = append(products, Product{id, productName, productType, supplierCost, retailPrice})

			fmt.Println("Successfully obtained products. Count of products = ", len(products))
		}

	} else {
		fmt.Println("WARNING. No results returned!")
	}

	fmt.Println("END GetProducts")
	return products
}
