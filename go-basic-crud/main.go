package main

import (
	"go-basic-crud/handler"
	"go-basic-crud/task"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	databaseString := "root:my-secret-pw@tcp(127.0.0.1:33060)/go-basic-crud?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(databaseString), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&task.Task{})

	taskRepository := task.NewRepository(db)

	taskService := task.NewService(taskRepository)

	taskHandler := handler.NewTaskHandler(taskService)

	router := gin.Default()

	api := router.Group("api")

	api.POST("/task", taskHandler.Store)
	api.GET("/task", taskHandler.Index)
	api.DELETE("/task/:id", taskHandler.Delete)

	router.Run()
}
