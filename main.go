package main

import (
	"fmt"
	"gin-gonic/controller" // import controller nama-projek/folder-controller
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func main() {

	r := gin.Default()
	r.POST("/signup", controller.SignUp)
	r.POST("/login", controller.Login)
	r.GET("/books", auth, controller.GetBooks)
	r.GET("/book", controller.GetBookDetail)
	r.GET("/trx", controller.GetTransaction)
	r.GET("/trx/detail", controller.GetTransactionDetail)

	r.Run(":8081") // listen and serve on 0.0.0.0:8080
}

func auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret-gin-gonic"), nil
	})
	fmt.Println(token)
	fmt.Println(err)

	if token != nil && err == nil {
		fmt.Println("token verified")
	} else {
		result := gin.H{
			"message": "not authorized",
			"error":   err.Error(),
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}
}
