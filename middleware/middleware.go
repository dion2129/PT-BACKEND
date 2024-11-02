package middleware

import (
    "net/http"
    "strings"
    "github.com/gin-gonic/gin"
    "api-test/helpers"
)

func AdminMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Mendapatkan token dari header Authorization
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
            c.Abort()
            return
        }

        token := strings.TrimPrefix(authHeader, "Bearer ")
        claims, err := helpers.ParseTokenJWT(token)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        // Memeriksa apakah role adalah "admin"
        role, ok := claims["role"].(string)
        if !ok || role != "admin" {
            c.JSON(http.StatusForbidden, gin.H{"error": "You do not have access to this resource"})
            c.Abort()
            return
        }

        c.Next()
    }
}
