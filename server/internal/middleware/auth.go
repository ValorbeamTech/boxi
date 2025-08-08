package middleware

import (
	"net/http"
	"strings"

	"github.com/francis/projectx-api/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, model.ErrorResponse("Authorization header required"))
            c.Abort()
            return
        }

        bearerToken := strings.Split(authHeader, " ")
        if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, model.ErrorResponse("Invalid authorization header format"))
            c.Abort()
            return
        }

        token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, jwt.ErrSignatureInvalid
            }
            return []byte(jwtSecret), nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, model.ErrorResponse("Invalid token"))
            c.Abort()
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            c.JSON(http.StatusUnauthorized, model.ErrorResponse("Invalid token claims"))
            c.Abort()
            return
        }

        userID, ok := claims["user_id"].(float64)
        if !ok {
            c.JSON(http.StatusUnauthorized, model.ErrorResponse("Invalid user ID in token"))
            c.Abort()
            return
        }

        c.Set("user_id", int(userID))
        c.Set("email", claims["email"])
        c.Next()
    }
}