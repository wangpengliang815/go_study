package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Replace with your own connection parameters
var server = "localhost"
var port = 1433
var user = "sa"
var database = "go_db"
var password = "wpl19950815"

var db *sql.DB

func main() {
	// 创建一个默认路由
	r := gin.Default()

	connStr := fmt.Sprintf("server=%s;port=%d;database=%s;user id=%s;password=%s", server, port, database, user, password)

	fmt.Printf("dbconn:%s", connStr)
	db, err := sql.Open("mssql", connStr)
	if err != nil {
		log.Fatal("creating connection error: " + err.Error())
	}
	log.Printf("Connected!\n")
	defer db.Close()
	// books := make([]int, 5)

	// getlist
	r.GET("/persons", func(c *gin.Context) {
		// db.QueryContext()
		// rows, err := db.Query("select * from Person")
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Println(rows)
		// defer rows.Close()
		// for rows.Next() {
		// 	err := rows.Scan(&id, &name)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	log.Println(id, name)
		// }
		c.JSON(http.StatusOK, 2)
	})

	// // post
	// r.POST("/books", func(c *gin.Context) {
	// 	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 	books = append(books, r.Intn(10))
	// 	c.JSON(http.StatusOK, books)
	// })

	// // put
	// r.PUT("/books", func(c *gin.Context) {
	// 	index := c.DefaultQuery("index", "0")
	// 	data, err := strconv.Atoi(index)
	// 	if err != nil {
	// 		return
	// 	}
	// 	books[data] = 999
	// 	c.JSON(http.StatusOK, books)
	// })

	// // delete
	// r.DELETE("/books", func(c *gin.Context) {
	// 	index := c.DefaultQuery("index", "0")
	// 	data, err := strconv.Atoi(index)
	// 	if err != nil {
	// 		return
	// 	}
	// 	books[data] = 999
	// 	books = append(books[:data], books[data:]...)
	// 	c.JSON(http.StatusOK, books)
	// })

	r.Run()
}
