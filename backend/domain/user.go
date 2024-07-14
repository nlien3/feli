package domain

import (
	"time"
)

// Usuario representa a un usuario en el sistema.
type Usuario struct {
	IdUsuario      int64     `gorm:"primaryKey;column:Id_usuario;autoIncrement" json:"id_usuario"`
	NombreUsuario  string    `gorm:"column:Nombre_usuario;not null" json:"nombre_usuario"`
	Contrasena     string    `gorm:"column:Contrasena;not null" json:"contrasena"`
	Tipo           string    `gorm:"column:Tipo;not null" json:"tipo"`
	CreatedAt      time.Time `gorm:"column:Created_at;autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:Updated_at;autoUpdateTime" json:"updated_at"`
}
