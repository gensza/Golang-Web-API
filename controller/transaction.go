package controller

import (
	"fmt"
	"gin-gonic/database"
	"gin-gonic/model" // import model nama-project/nama-folder model
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTransaction(c *gin.Context) {

	var t model.Transaction
	var ts []model.Transaction

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
	var query = "SELECT id, no_ref, id_book, qty, transaction_at FROM tb_transaction WHERE 1=1"

	fmt.Println(query)

	rows, err := database.InitDB().Query(query+" ORDER BY id DESC LIMIT ?,?", pageOffset, limit)

	defer rows.Close()

	if err != nil {
		fmt.Println("DB Query : ", err.Error())
	}

	for rows.Next() {
		err = rows.Scan(&t.Id, &t.Noref, &t.IdBook, &t.Qty, &t.Transaction_at)
		if err != nil {
			fmt.Println("Scan :", err.Error())
		}
		ts = append(ts, t)
	}

	if ts != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"status":  http.StatusOK,
			"data":    ts,
		})
	} else {
		c.JSON(http.StatusNoContent, gin.H{
			"message": "No content available",
			"status":  http.StatusNoContent,
			"data":    "",
		})
	}
}

func GetTransactionDetail(c *gin.Context) {

	var t model.Transaction
	var ts []model.Transaction
	var cr model.Courier

	IdTrx := c.Query("id_trx")
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

	// query tb_trx
	var query = "SELECT id, no_ref, id_book, qty, transaction_at FROM tb_transaction WHERE 1=1 AND id = ?"

	fmt.Println(query)

	rows, err := database.InitDB().Query(query+" ORDER BY id DESC LIMIT ?,?", IdTrx, pageOffset, limit)

	defer rows.Close()

	if err != nil {
		fmt.Println("DB Query : ", err.Error())
	}

	for rows.Next() {
		err = rows.Scan(&t.Id, &t.Noref, &t.IdBook, &t.Qty, &t.Transaction_at)
		if err != nil {
			fmt.Println("Scan :", err.Error())
		}
		// ts = append(ts, t)
	}

	// query tb_courier
	var query2 = "SELECT a.id, a.id_user, b.name FROM tb_courier a LEFT JOIN tb_user b ON a.id_user = b.id WHERE a.id_transaction = ?"

	fmt.Println(query2)

	rows2, err2 := database.InitDB().Query(query2, t.Id)

	defer rows2.Close()

	if err2 != nil {
		fmt.Println("DB Query : ", err2.Error())
	}

	for rows2.Next() {
		err2 = rows2.Scan(&cr.IdCourier, &cr.IdUser, &cr.Name)
		if err != nil {
			fmt.Println("Scan :", err2.Error())
		}
		t.Courier = append(t.Courier, cr)
	}

	ts = append(ts, t)

	if t.Id != 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"status":  http.StatusOK,
			"data":    t,
		})
	} else {
		c.JSON(http.StatusNoContent, gin.H{
			"message": "No content available",
			"status":  http.StatusNoContent,
			"data":    "",
		})
	}
}
