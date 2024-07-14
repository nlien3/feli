package controllers

import (
	"backend/clients"
	"backend/dao"
	"backend/domain"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateChatMessage(c *gin.Context) {
	var params dao.Chat

        // Vincula el cuerpo de la solicitud a la estructura
        if err := c.ShouldBindJSON(&params); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // Crear una instancia de la estructura de la base de datos
        param := dao.Chat{
            IdUsuario: int(params.IdUsuario),
            IdCurso: int(params.IdCurso),
			Message: params.Message,
        }


		// Verificar que el curso existe
    	var course dao.Course
    	if err := clients.DB.First(&course, param.IdCurso).Error; err != nil {
       	 	c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
       	 	return
    	}

    	// Verificar que el usuario existe
   		var user dao.Usuario
    	if err := clients.DB.First(&user, param.IdUsuario).Error; err != nil {
        	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        	return
    	}

        // Guardar los datos en la base de datos
        if err := clients.DB.Create(&param).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        // Responder con los datos guardados
        c.JSON(http.StatusOK, param)

}


func GetChatMessagesByCourseID(c *gin.Context) {
	var chatMessages []dao.Chat
    var response []domain.ChatResponse
    courseId := c.Param("courseId")

    // Verificar que el curso existe
    var course dao.Course
    if err := clients.DB.First(&course, courseId).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "The specified course does not exist"})
        return
    }

    // Obtener los mensajes de chat ordenados por CreatedAt en orden descendente
    if err := clients.DB.Where("Id_curso = ?", courseId).Order("Created_at desc").Find(&chatMessages).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Iterar sobre los mensajes de chat y obtener los datos del usuario
    for _, chat := range chatMessages {
        var user dao.Usuario
        if err := clients.DB.First(&user, chat.IdUsuario).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        response = append(response, domain.ChatResponse{
            IdChat:        chat.IdChat,
            IdUsuario:     chat.IdUsuario,
            NombreUsuario: user.NombreUsuario,
            IdCurso:       chat.IdCurso,
            Message:       chat.Message,
            CreatedAt:     chat.CreatedAt,
            UpdatedAt:     chat.UpdatedAt,
        })
    }

    if len(chatMessages) == 0 {
        c.JSON(http.StatusOK, []domain.ChatResponse{})
        return
    }

    c.JSON(http.StatusOK, response)
}


func UploadFileHandler(c *gin.Context) {
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
        return
    }

    // Guardar el archivo con un nombre único
    filename := filepath.Base(file.Filename)
    newFilename := filepath.Join("uploads", time.Now().Format("20060102150405")+"_"+filename)

    if err := c.SaveUploadedFile(file, newFilename); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save the file"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "filename": newFilename})
}

