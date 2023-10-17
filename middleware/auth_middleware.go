package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saepudinasep/task-5-pbi-btpns-AsepSaepudin/helpers"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is missing"})
			c.Abort()
			return
		}

		claims, err := helpers.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Disimpan dalam konteks pengguna untuk penggunaan lebih lanjut
		c.Set("userID", claims.UserID)
		c.Next()
	}
}
