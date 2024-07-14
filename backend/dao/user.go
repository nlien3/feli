package dao

import (
	"time"
)

type Usuario struct {
	IdUsuario     int       `gorm:"primaryKey;column:Id_usuario;autoIncrement"`
	NombreUsuario string    `gorm:"column:Nombre_usuario;not null;unique"`
	Contrasena    string    `gorm:"column:Contrasena;not null"`
	Tipo          string    `gorm:"column:Tipo;not null"` // admin or normal
	CreatedAt     time.Time `gorm:"column:Created_at;autoCreateTime"`
	UpdatedAt     time.Time `gorm:"column:Updated_at;autoUpdateTime"`
}

type LoginRequest struct {
	NombreUsuario string `json:"nombre_usuario"`
	Contrasena    string `json:"contrasena"`
}
