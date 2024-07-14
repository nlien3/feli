package dao

import (
	"time"
)

type Course struct {
	ID          	uint      `gorm:"primaryKey"`
	Nombre      	string    `json:"nombre"`
	Categoria		string 	  `json:"categoria"` 
	Dificultad 	 	string    `json:"dificultad"`
	Precio      	float64   `json:"precio"`
	Descripcion   	string    `json:"descripcion"`
	ImageURL    	string    `json:"imageURL"`
	CreatedAt   	time.Time `json:"created_at"`
	UpdatedAt   	time.Time `json:"updated_at"`
}

