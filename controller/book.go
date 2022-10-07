package controller

import (
	"fmt"
	"gin-gonic/database" // import database nama-project/nama-folder database
	"gin-gonic/model"    // import model nama-project/nama-folder model
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {

	offset := c.Query("offset")
	page := c.Query("page")
	limit := c.Query("limit")
	rating := c.Query("rating")

	var pageOffset int
	var pageInt, _ = strconv.Atoi(page)
	var limitInt, _ = strconv.Atoi(limit)

	if offset != "" && page == "" {
		pageOffset, _ = strconv.Atoi(offset)
	}
	if limit == "" {
		limitInt = 10
	}
	if page != "" {
		pageOffset = (limitInt * pageInt) - limitInt
	}

	var data model.BookModel

	data.Rating = rating

	result := model.GetBooksModel(data, pageOffset, limit)

	// fmt.Println(err)

	if result != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"status":  http.StatusOK,
			"data":    result,
		})
	} else {
		c.JSON(http.StatusNoContent, gin.H{
			"message": "No content available",
			"status":  http.StatusNoContent,
			"data":    "",
		})
	}
}

func GetBookDetail(c *gin.Context) {

	idBook := c.Query("id_book")

	var book model.BookModel
	var books []model.BookModel

	rows, err := database.InitDB().Query("SELECT id, title, description, price, rating, discount FROM books where id = ?", idBook)
	if err != nil {
		fmt.Println("DB Query : ", err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&book.Id, &book.Title, &book.Description, &book.Price, &book.Rating, &book.Discount)
		if err != nil {
			fmt.Println("Scan :", err.Error())
		}
		books = append(books, book)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"status":  http.StatusOK,
		"data":    book,
	})
}
