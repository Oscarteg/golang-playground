package main

import (
	"go-basic-crud/handler"
	"go-basic-crud/task"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/swaggo/gin-swagger/example/basic/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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
	api.GET("/task/:id", taskHandler.Find)
	api.PUT("/task/:id", taskHandler.Update)
	api.DELETE("/task/:id", taskHandler.Delete)


	// swagger
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run()
}
