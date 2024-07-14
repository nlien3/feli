package domain

import "time"

type ChatResponse struct {
    IdChat        int       `json:"IdChat"`
    IdUsuario     int       `json:"IdUsuario"`
    NombreUsuario string    `json:"NombreUsuario"`
    IdCurso       int       `json:"IdCurso"`
    Message       string    `json:"Message"`
    CreatedAt     time.Time `json:"CreatedAt"`
    UpdatedAt     time.Time `json:"UpdatedAt"`
}
