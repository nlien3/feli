package controllers

import (
	"backend/clients"
	"backend/dao"
	"backend/domain"
	"backend/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateSubscription(c *gin.Context) {
    var request domain.SubscribeRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        log.Printf("Error binding JSON: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
        return
    }

    // Logging the received request for debugging
    log.Printf("Received subscription request: %+v\n", request)

    userIDJWT, exists := c.Get("userID")
    if !exists {
     c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
     return
    }

    // Aserción de tipo para convertir userIDJWT a int
    userIDJWTINT, ok := userIDJWT.(int)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
        return
    }

    // Validate IDs are not zero
    if request.UserID == 0 || request.CourseID == 0 {
        log.Printf("Invalid user ID or course ID: UserID=%d, CourseID=%d", userIDJWTINT, request.CourseID)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID or course ID"})
        return
    }

    subscription := dao.Subscription{
        IdUsuario: userIDJWTINT,
        IdCurso:   int(request.CourseID),
    }

    log.Printf("Creating subscription with User ID: %d and Course ID: %d", subscription.IdUsuario, subscription.IdCurso)
    newSubscription, err := services.CreateSubscription(subscription)
    if err != nil {
        log.Printf("Error creating subscription: %v", err)
        if err.Error() == "user not found" || err.Error() == "course not found" {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating subscription", "details": err.Error()})
        }
        return
    }
    log.Println("Subscription created successfully")
    c.JSON(http.StatusCreated, newSubscription)
}

func GetSubscriptions(c *gin.Context) {
	var subscriptions []dao.Subscription
	if err := clients.DB.Find(&subscriptions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching subscriptions"})
		return
	}
	c.JSON(http.StatusOK, subscriptions)
}

func DeleteSubscription(c *gin.Context) {
	var subscription dao.Subscription
	id := c.Param("id")
	if err := clients.DB.First(&subscription, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subscription not found"})
		return
	}
	if err := clients.DB.Delete(&subscription).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting subscription"})
		return
	}
	c.Status(http.StatusNoContent)
}

func GetSubscriptionsByUser(c *gin.Context) {
    idUserStr := c.Param("iduser")
    log.Println("Tipo de variable",idUserStr)
    idUser, err := strconv.Atoi(idUserStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    userIDJWT, exists := c.Get("userID")
    if !exists {
     c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
     return
    }

    // Aserción de tipo para convertir userIDJWT a int
    userIDJWTINT, ok := userIDJWT.(int)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
        return
    }

    // Verifica que el usuario exista en caso contrario da error
    var user dao.Usuario
    resultExist := clients.DB.First(&user, userIDJWTINT)
        if resultExist.Error != nil {
            if resultExist.Error == gorm.ErrRecordNotFound {
                c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no existe"})
            } else {
                c.JSON(http.StatusInternalServerError, gin.H{"error": resultExist.Error.Error()})
            }
            return
        }

    // Busca las suscripciones por id del usuario
    var subscriptions []dao.Subscription
    result := clients.DB.Where("Id_usuario = ?", idUser).Find(&subscriptions)

    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

     // Obtener los datos del curso para cada suscripción
        type SubscriptionWithCourse struct {
            dao.Subscription
            Curso dao.Course `json:"curso"`
        }

    var response []SubscriptionWithCourse
        for _, subscription := range subscriptions {
            var curso dao.Course
            resultCurso := clients.DB.First(&curso, subscription.IdCurso)
            if resultCurso.Error != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": resultCurso.Error.Error()})
                return
            }

            response = append(response, SubscriptionWithCourse{
                Subscription: subscription,
                Curso:        curso,
            })

        }
          if len(subscriptions) == 0 {
            c.JSON(http.StatusOK, []dao.Subscription{})
            return
        }
        c.JSON(http.StatusOK, response)
}