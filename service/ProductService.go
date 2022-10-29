package service

import (
	"TimeMachineApi/database"
	"fmt"
)

type Product struct {
	Id           int    `json:"id"`
	ProductName  string `json:"productName"`
	ProductType  string `json:"productType"`
	SupplierCost string `json:"supplierCost"`
	SalePrice    string `json:"salePrice"`
}

func GetProducts() []Product {

	var products []Product
	fmt.Println("BEGIN GetProducts")

	db := database.OpenConnection()

	if db != nil {

		rows, _ := db.Query("SELECT id, productName, productType, supplierCost, salePrice FROM timemachine.product")

		if rows != nil {

			for rows.Next() {
				var id int
				var productName string
				var productType string
				var supplierCost string
				var retailPrice string

				err := rows.Scan(&id, &productName, &productType, &supplierCost, &retailPrice)

				if err != nil {
					fmt.Println("Error trying to scan row for the listed fields: ", err)
					database.CloseConnection(db)
					return nil
				}

				products = append(products, Product{id, productName, productType, supplierCost, retailPrice})

				fmt.Println("Successfully obtained products. Count of products = ", len(products))
			}

		} else {
			fmt.Println("WARNING. No results returned!")
			database.CloseConnection(db)
		}
	}

	fmt.Println("END GetProducts")
	return products
}
