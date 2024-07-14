// controllers/users_controller.go
package controllers

import (
	"backend/dao"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var registerRequest struct {
		NombreUsuario string `json:"nombre_usuario"`
		Contrasena    string `json:"contrasena"`
		Tipo          string `json:"tipo"`
	}

	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.RegisterS(registerRequest.NombreUsuario, registerRequest.Contrasena, registerRequest.Tipo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var loginRequest dao.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, userID, userType, err := services.Login(loginRequest.NombreUsuario, loginRequest.Contrasena)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":  token,
		"userId": userID,
		"type":   userType,
	})
}


func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

func ProtectedEndpoint(c *gin.Context) {
	username, _ := c.Get("username")
	c.JSON(http.StatusOK, gin.H{"message": "Hello, " + username.(string)})
}
