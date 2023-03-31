package controller

import (
	"Goal-Back-End/database"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Patient struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	OrderId int    `json:"orderId"`
}

type Message struct {
	Message string `json:"message"`
}

type Order struct {
	OrderId int    `json:"orderId"`
	Message string `json:"message"`
}

func GetList(c *gin.Context) {
	rows, err := database.SqlDB.Query("SELECT name, id, order_id FROM patient ORDER BY id")

	// Handle errors
	if err != nil {
		log.Fatal(err)
	}

	// Create an empty slice to hold the results
	var patients []Patient

	// Loop through the rows of data and append to the slice
	for rows.Next() {
		var p Patient
		err := rows.Scan(&p.Name, &p.Id, &p.OrderId)
		if err != nil {
			log.Fatal(err)
		}
		patients = append(patients, p)
	}

	// Return the data as JSON
	c.JSON(http.StatusOK, patients)
}

func GetOrderById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	row := database.SqlDB.QueryRow("SELECT message FROM patient WHERE order_id=$1", id)

	var message string

	err = row.Scan(&message)
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "order not found"})
			return
		} else {
			panic(err)
		}
	}

	// Create a map with the message data
	messageData := map[string]interface{}{
		"message": message,
	}

	// Return the message data as JSON
	c.JSON(http.StatusOK, messageData)
}

func UpdateOrderById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	// Parse the JSON data from the request body
	var reqBody Message
	if err := c.BindJSON(&reqBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	result, err := database.SqlDB.Exec("UPDATE patient SET message = $1 WHERE order_id = $2", reqBody.Message, id)
	if err != nil {
		panic(err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data updated successfully"})
}
