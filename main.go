package main

import (
	"asset-tracker/database"
	"asset-tracker/middleware"
	"asset-tracker/usecase"
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	var DB *sql.DB
	var err error

	// env config
	err = godotenv.Load(".env")
	if err != nil {
		fmt.Println("error when load environment file")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	DB, _ = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connection failed")
		panic(err)
	}

	database.DbMigrate(DB)

	defer DB.Close()

	router := gin.Default()

	user := router.Group("/user")
	user.POST("/register", usecase.RegisterUser)
	user.POST("/login", usecase.LoginUser)

	user.PUT("/", middleware.AuthUser(), usecase.UpdateProfile)

	router.Run(os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT"))
}
