package controller

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	// var user model.Login
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username != "myname" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "wrong username",
		})
	} else {
		if password != "myname123" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "wrong password",
			})
		} else {

			sign := jwt.New(jwt.GetSigningMethod("HS256"))
			claims := sign.Claims.(jwt.MapClaims)
			claims["username"] = username
			claims["email"] = "gensza@gmail.com"
			claims["exp"] = time.Now().Add(time.Hour * 168).Unix()
			token, err := sign.SignedString([]byte("secret-gin-gonic"))

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
				c.Abort()
			}
			c.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		}
	}
}
