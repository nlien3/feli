package domain

// Subscription representa una suscripción de un usuario a un curso.
// Esto no deberia ir porque esta en dao.
/*type Subscription struct {
	IdSubscription int64     `gorm:"primaryKey;column:Id_subscription;autoIncrement" json:"id_subscription"`
	IdUsuario      int64     `gorm:"column:Id_usuario;not null" json:"id_usuario"`
	IdCurso        int64     `gorm:"column:Id_curso;not null" json:"id_curso"`
	CreatedAt      time.Time `gorm:"column:Created_at;autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:Updated_at;autoUpdateTime" json:"updated_at"`
}*/

// SubscribeRequest represents the payload for subscribing to a course.
type SubscribeRequest struct {
	UserID   int64 `json:"userID"`
	CourseID int64 `json:"courseID"`
}