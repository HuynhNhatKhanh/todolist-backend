package config

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CorsConfig returns a gin.HandlerFunc which enables Cross-Origin Resource
// Sharing (CORS) for a Gin engine. The default config allows GET, POST, PUT,
// DELETE, OPTIONS methods from http://localhost:5173 and https://yourdomain.com
// with Content-Type, Authorization, X-Requested-With headers and credentials.
// The preflight request is cached for 12 hours.
func CorsConfig() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://todolist-frontend-ten.vercel.app/"},
		// AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "X-Requested-With"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
