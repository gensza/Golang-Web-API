package controller

import (
	"fmt"
	"gin-gonic/database"
	"gin-gonic/model"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var bind model.Login

	username := c.PostForm("username")
	password := c.PostForm("password")

	cekUser := database.InitDB().QueryRow("SELECT username, password FROM tb_user WHERE username = ?", username).Scan(&bind.Username, &bind.Password)
	cekPass := bcrypt.CompareHashAndPassword([]byte(bind.Password), []byte(password))

	if cekUser != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "wrong username",
		})
	} else {
		if cekPass != nil {
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

func SignUp(c *gin.Context) {
	var bind model.SignUp

	// validation data input
	err := c.ShouldBindJSON(&bind)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on Field %s, condition: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": errorMessage,
			})
			return
		}
	}

	// cek username existing
	rows, err := database.InitDB().Query("SELECT COUNT(username) FROM tb_user WHERE username = ?", bind.Username)
	if err != nil {
		fmt.Println("DB Query Fail:" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "DB Query Fail :" + err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}
	defer rows.Close()

	var countUser int
	for rows.Next() {
		err = rows.Scan(&countUser)
		if err != nil {
			fmt.Println("DB Scan Fail :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "DB Scan Fail:" + err.Error(),
				"status":  http.StatusInternalServerError,
			})
		}
	}

	if countUser >= 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Username Already Existing",
			"status":  http.StatusBadRequest,
		})
		return
	}

	// hash password and compare hash password
	hashPassword, _ := HashPassword(bind.Password)
	matchPassword := CheckPasswordHash(bind.Password, hashPassword)

	if !matchPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Hash Password Fail",
			"status":  http.StatusBadRequest,
		})
		return
	}

	// prepare query Insert
	stmt, err := database.InitDB().Prepare("INSERT INTO tb_user(name, address, username, password) VALUES(?,?,?,?)")
	if err != nil {
		fmt.Println("DB Prepare Fail : ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "DB Prepare Fail:" + err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}
	// Exec queery insert
	_, err = stmt.Exec(bind.Name, bind.Address, bind.Username, hashPassword)
	if err != nil {
		fmt.Println("Query Exec Fail : ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "DB Prepare Exec:" + err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}
	defer stmt.Close()

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"status":  http.StatusCreated,
	})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
