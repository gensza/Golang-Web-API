package model

import (
	"fmt"
	"gin-gonic/database"
)

type ProductModel struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Price int    `json:"price"`
}

func GetProductsModel(pageOffset int, limit string) []ProductModel {

	// var product
	var product ProductModel
	var products []ProductModel

	var query = "SELECT id, title, price FROM products"

	rows, err := database.InitDB().Query(query)
	if err != nil {
		fmt.Println("DB Query : ", err.Error())
	}

	for rows.Next() {
		err = rows.Scan(&product.Id, &product.Title, &product.Price)
		if err != nil {
			fmt.Println("Scan :", err.Error())
		}
		products = append(products, product)
	}
	defer rows.Close()

	return products
}

func DelProductsModel(id int) interface{} {

	var query = "DELETE FROM products WHERE id =?"

	result, err := database.InitDB().Query(query, id)
	if err != nil {
		fmt.Println("DB Query : ", err.Error())
	}
	defer result.Close()

	return err
}
