package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rlevidev/oauth-go/src/routes"
)

func main() {
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)
	
	router.Run()
}