package main

import (
	"gin-gonic/controller" // import controller nama-projek/folder-controller

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/books", controller.GetBooks)
	r.GET("/book", controller.GetBookDetail)
	r.GET("/trx", controller.GetTransaction)
	r.GET("/trx/detail", controller.GetTransactionDetail)

	r.Run(":8081") // listen and serve on 0.0.0.0:8080
}
