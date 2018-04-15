package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type ToDo struct {
	gorm.Model
	Title string
	State bool
}

func main() {
	dbDSN := fmt.Sprintf("user=%s password=%s DB.name=gtd port=5432 sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))
	db, err := gorm.Open("postgres", dbDSN)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&ToDo{})
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/todos", func(c *gin.Context) {
		var todo ToDo
		c.BindJSON(&todo)

		db.Create(&todo)
	})

	r.GET("/todos", func(c *gin.Context) {
		var todos []ToDo
		db.Find(&todos)
		c.JSON(200, &todos)
	})

	r.Run()
}
