package domain

import "time"

type Course struct {
	IdCurso    		int64     `json:"id"`          // Course ID
	Nombre     		string    `json:"title"`       // Course title
	Categoria		string 	  `json:"categoria"`
	Dificultad 		string    `json:"description"` // Course description
	Precio     		string    `json:"category"`    // Course Category. Allowed values: to be defined
	Descripcion		string    `json:"descripcion"`
	ImageURL   		string    `json:"image_url"`
	CreatedAt  		time.Time `json:"creation_date"` // Course creation date
	UpdatedAt  		time.Time `json:"last_updated"`  // Course last updated date
}

type SearchResponse struct {
	Results []Course `json:"results"`
}

type SearchRequest struct {
	IdCurso int `json:"IdCurso"`
}

