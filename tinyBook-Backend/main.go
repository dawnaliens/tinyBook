package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	"tinyBook/tinyBook-Backend/internal/web"
)

func main() {
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		// AllowOrigins:  []string{"https://localhost:3000"},
		//AllowMethods:  []string{"POST"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
		// ExposeHeaders: []string{"Content-Type", "Authorization"},
		// Credentials - ask for bringing cookies
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				// Development environment
				return true
			}
			return strings.Contains(origin, "testing.com")
		},
		MaxAge: 12 * time.Hour,
	}))

	u := web.NewUserHandler()
	u.RegisterRoutes(server)
	server.Run(":8080")
}
