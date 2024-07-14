package services

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"backend/clients"
	"backend/dao"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	UserID   int    `json:"userId"`
	jwt.StandardClaims
}

func GenerateJWT(username string, userID int) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		UserID:   userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func md5Hash(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Login(username, password string) (string, int, string, error) {
	fmt.Println("Username:", username) // Depuración

	// Hashear la contraseña recibida del frontend usando MD5
	hashedPassword := md5Hash(password)

	fmt.Println("Hashed Password from Frontend (MD5):", hashedPassword) // Depuración

	var user dao.Usuario
	// Buscar usuario por nombre de usuario
	if err := clients.DB.Where("nombre_usuario = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("User not found") // Depuración
			return "", 0, "", errors.New("invalid credentials")
		}
		return "", 0, "", err
	}

	fmt.Println("User found:", user.NombreUsuario) // Depuración
	fmt.Println("Hashed Password in DB:", user.Contrasena) // Depuración

	// Comparar la contraseña hasheada del frontend (MD5) con la almacenada en la base de datos
	if user.Contrasena != hashedPassword {
		fmt.Println("Invalid password") // Depuración
		return "", 0, "", errors.New("invalid credentials")
	}

	// Generar token JWT si la contraseña es correcta
	token, err := GenerateJWT(username, user.IdUsuario)
	if err != nil {
		return "", 0, "", err
	}

	return token, user.IdUsuario, user.Tipo, nil
}

func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, err
		}
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
