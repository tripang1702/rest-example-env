package main

import (
	"log"
	"os"
	"rest-example-env/controller"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	ctr := controller.NewController()

	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	database := os.Getenv("DB_NAME")

	ctr.SetParamEnv(server, port, user, password, database)
	// end

	gin.SetMode(gin.ReleaseMode) // set this to production
	router := gin.Default()
	router.SetTrustedProxies(nil)

	scdb := router.Group("api")
	{
		scdb.GET("example", ctr.GetExample)

	}

	router.Run(":8090")
}
