package main

import (
	"github.com/gin-gonic/gin"
	"tinyBook/tinyBook-Backend/internal/web"
)

func main() {
	server := gin.Default()

	u := web.NewUserHandler()
	u.RegisterRoutes(server)
	server.Run(":8080")
}
