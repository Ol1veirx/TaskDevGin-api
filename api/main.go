package main

import (
	"TaskDevGin-api/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}
	ConnectDatabase()
	DB.AutoMigrate(&handlers.Task{})

	router := gin.Default()

	router.GET("/api/v1/tasks", handlers.GetTasks)
	router.GET("/api/v1/tasks/:id", handlers.GetTaskById)
	router.POST("/api/v1/tasks", handlers.CreateTask)
	router.PUT("/api/v1/tasks/:id", handlers.UpdateTask)
	router.DELETE("/api/v1/tasks/:id", handlers.DeleteTask)

	router.Run(":8080")
}

var DB *gorm.DB

func ConnectDatabase() {
	requiredEnvVars := []string{"DB_HOST", "DB_USER", "DB_PASSWORD", "DB_NAME", "DB_PORT"}
	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			log.Fatalf("A variável de ambiente %s não está definida", envVar)
		}
	}

	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
	handlers.DB = db
}
