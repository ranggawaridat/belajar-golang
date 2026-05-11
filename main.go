package main

import (
	"github.com/ranggawaridat/belajar-golang/database"
	"github.com/ranggawaridat/belajar-golang/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
}
