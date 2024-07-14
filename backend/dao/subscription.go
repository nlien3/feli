package dao

import (
	"time"
)

type Subscription struct {
	IdSubscription int       `gorm:"primaryKey;column:Id_subscription;autoIncrement"`
	IdUsuario      int       `gorm:"column:Id_usuario;not null"`
	IdCurso        int       `gorm:"column:Id_curso;not null"`
	CreatedAt      time.Time `gorm:"column:Created_at;autoCreateTime"`
	UpdatedAt      time.Time `gorm:"column:Updated_at;autoUpdateTime"`
}
