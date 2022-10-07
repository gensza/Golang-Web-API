package model

import (
	"fmt"
	"gin-gonic/database"
)

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignUp struct {
	Name     string `json:"name" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func GetBooksModel(data BookModel, pageOffset int, limit string) []BookModel {

	var book BookModel
	var books []BookModel

	var query = "SELECT id, title, description, price, rating, discount FROM books WHERE 1=1"

	if data.Rating != "" {
		query = query + " AND rating = " + data.Rating
	}

	rows, err := database.InitDB().Query(query+" ORDER BY id DESC LIMIT ?,?", pageOffset, limit)
	if err != nil {
		fmt.Println("DB Query : ", err.Error())
	}

	for rows.Next() {
		err = rows.Scan(&book.Id, &book.Title, &book.Description, &book.Price, &book.Rating, &book.Discount)
		if err != nil {
			fmt.Println("Scan :", err.Error())
		}
		books = append(books, book)
	}
	defer rows.Close()

	return books
}
