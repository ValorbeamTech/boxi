package middleware

import (
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
    config := cors.DefaultConfig()
    config.AllowOrigins = []string{"*"} // Configure for production
    config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
    config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
    
    return cors.New(config)
}