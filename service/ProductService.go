package service

import (
	"TimeMachineApi/database"
	log "github.com/sirupsen/logrus"
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
	log.Info("BEGIN GetProducts")

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
					log.Error("Error trying to scan row for the listed fields: ", err)
					database.CloseConnection(db)
					return nil
				}

				var p = Product{id, productName, productType, supplierCost, retailPrice}
				products = append(products, p)

				log.Debug("Product queried and added to the list of products = ", p)
				log.Info("Successfully obtained products. Count of products = ", len(products))
			}

		} else {
			log.Warn("WARNING. No results returned!")
			database.CloseConnection(db)
		}
	}

	log.Info("END GetProducts")
	return products
}
