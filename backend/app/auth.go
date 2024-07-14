// app/auth.go
package app

import (
	"backend/clients"
	"backend/dao"
	"backend/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthNormal() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "no token provided"})
			c.Abort()
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]
		claims, err := services.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

   		var usuario dao.Usuario
        clients.DB.Where("Nombre_usuario = ?", claims.Username).First(&usuario)
        if usuario.Tipo != "normal" {
            c.JSON(http.StatusForbidden, gin.H{"error": "No tienes el rol válido"})
            c.Abort()
            return
		}
		c.Set("userID", claims.UserID)
		c.Next()
	}
}

func AuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "no token provided"})
			c.Abort()
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]
		claims, err := services.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

   		var usuario dao.Usuario
        clients.DB.Where("Nombre_usuario = ?", claims.Username).First(&usuario)
        if usuario.Tipo != "admin" {
            c.JSON(http.StatusForbidden, gin.H{"error": "No tienes el rol válido"})
            c.Abort()
            return
		}
		c.Set("userID", claims.UserID)
		c.Next()
	}
}

func AuthAmbos() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "no token provided"})
			c.Abort()
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]
		claims, err := services.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

   		
		c.Set("userID", claims.UserID)
		c.Next()
	}
}
