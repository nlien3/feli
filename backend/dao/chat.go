package dao

import (
	"time"
)

type Chat struct {
    IdChat     		int       `gorm:"primaryKey;column:Id_chat;autoIncrement"`
    IdUsuario  		int       `gorm:"column:Id_usuario;not null"`
    IdCurso    		int       `gorm:"column:Id_curso;not null"`
    Message    		string    `gorm:"column:Message;type:text;not null"`
    CreatedAt  		time.Time `gorm:"column:Created_at;autoCreateTime"`
    UpdatedAt  		time.Time `gorm:"column:Updated_at;autoUpdateTime"`
}