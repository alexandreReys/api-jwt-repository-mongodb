package main

import (
	"log"

	"github.com/alexandreReys/api-jwt-repository-mongodb/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

	router := gin.Default();
	router.RouterGroup = *router.Group("/api/v1")	
	routes.InitRoutes(&router.RouterGroup)
	if err = router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
