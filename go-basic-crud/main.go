package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-basic-crud/handler"
	"go-basic-crud/task"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
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
// @BasePath /api
// @query.collection.format multi
// @in header
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
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://google.com"}
	// config.AllowOrigins == []string{"http://google.com", "http://facebook.com"}

	router.Use(cors.New(config))

	//router.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"http://localhost:8088"},
	//	AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
	//	AllowHeaders:     []string{"Origin", "Authorization", "x-api-key", "content-type"},
	//	ExposeHeaders:    []string{"Content-Length"},
	//	AllowCredentials: true,
	//	MaxAge: 12 * time.Hour,
	//}))

	api := router.Group("/api")
	{
		taskApi := api.Group("/tasks")
		{
			taskApi.POST("", taskHandler.Store)
			taskApi.GET("", taskHandler.Index)
			taskApi.GET("/:id", taskHandler.Find)
			taskApi.PUT("/:id", taskHandler.Update)
			taskApi.DELETE("/:id", taskHandler.Delete)
		}

	}

	// swagger
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run()
}
