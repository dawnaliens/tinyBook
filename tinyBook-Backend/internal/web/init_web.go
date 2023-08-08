package web

import "github.com/gin-gonic/gin"

func RegisterRoutes() *gin.Engine {
	server := gin.Default()

	u := &UserHandler{}
	u.RegisterRoutes(server)

	return server
}
