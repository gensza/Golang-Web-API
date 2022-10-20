package controller

import (
	"gin-gonic/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProduct(c *gin.Context) {

	offset := c.Query("offset")
	page := c.Query("page")
	limit := c.Query("limit")

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

	result := model.GetProductsModel(pageOffset, limit)

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

func DelProduct(c *gin.Context) {

	var id, _ = strconv.Atoi(c.Query("id"))

	err := model.DelProductsModel(id)

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"status":  http.StatusOK,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"status":  http.StatusBadRequest,
		})
	}
}

func InsertProduct(c *gin.Context) {

	var title = c.PostForm("title")
	var price = c.PostForm("price")

	var data model.ProductModel

	data.Title = title
	data.Price, _ = strconv.Atoi(price)

	err := model.InsertProductsModel(data)

	if err == nil {
		c.JSON(http.StatusCreated, gin.H{
			"message": "success",
			"status":  http.StatusCreated,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"status":  http.StatusBadRequest,
		})
	}
}
