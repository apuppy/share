package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Pong ping pong test
func Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// DbDemo go mysql
func DbDemo(c *gin.Context) {
	// "user:password@tcp/dbname?charset=utf8mb4"
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8mb4")
	// db, err := sql.Open("mysql", "user7:s$cret@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Prepare statement for inserting data
	insertSQL := "INSERT INTO `blog`.`link`(`title`, `url`, `is_delete`, `created_at`, `updated_at`, `deleted_at`) VALUES ('yu_tian_jian', 'http://yutianjian.com/', 0, '2020-11-02 11:23:01', NULL, NULL);"
	stmtIns, err := db.Prepare(insertSQL) // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates
	_, err = stmtIns.Exec()
	if err != nil {
		log.Fatal(err)
	}
}

// Category category save struct
type Category struct {
	Name string `form:"category_name" json:"category_name" xml:"category_name"  binding:"required"`
}

// SaveCategory save link category
func SaveCategory(c *gin.Context) {
	var json Category
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := saveCategoryToDB(json.Name); err != nil {
		Error(c, err.Error())
	}
	Success(c, nil)
}

func saveCategoryToDB(categoryName string) error {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8mb4")
	if err != nil {
		return err
	}
	defer db.Close()

	sql := "INSERT INTO link_category (name,sort,created_at) VALUES (?,?,?);"
	stmtIns, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	_, err = stmtIns.Exec(categoryName, 0, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		return err
	}
	return nil
}
