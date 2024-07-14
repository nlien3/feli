package app

import (
	"backend/controllers"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Middlewares
	router.Use(AllowCors())

	// Rutas y controladores
	router.GET("/courses", controllers.GetCourses)
	router.GET("/courses/:id", controllers.SearchByID)
	router.GET("/courses/name/:name", controllers.SearchByName)	
	router.GET("/courses/category/:category", controllers.GetCoursesByCategory)
	
	router.POST("/login", controllers.Login)
	router.POST("/register", controllers.RegisterC)
	
	// Ruta para subir archivos
    router.POST("/upload", controllers.UploadFileHandler)


	// Rutas protegidas
	protectedAdmin := router.Group("/").Use(AuthAdmin())
	protectedNormal := router.Group("/").Use(AuthNormal())
	protectedAmbos := router.Group("/").Use(AuthAmbos())

	
	{
		protectedAdmin.POST("/courses", controllers.CreateCourse)
		protectedAdmin.DELETE("/courses/:id", controllers.DeleteCourse)
		protectedAdmin.PUT("/courses/:id", controllers.UpdateCourseHandler)
	}

	{
		protectedNormal.GET("/subscriptions", controllers.GetSubscriptions)
		protectedNormal.POST("/subscriptions", controllers.CreateSubscription)
		protectedNormal.DELETE("/subscriptions/:id", controllers.DeleteSubscription)
		protectedNormal.GET("/subscriptions/:iduser", controllers.GetSubscriptionsByUser)
		
	}

	{
		protectedAmbos.GET("/courses/suscription/:id", controllers.GetSuscriptionByIdUser)
		protectedAmbos.POST("/course/chat", controllers.CreateChatMessage)
		protectedAmbos.GET("/course/:courseId/chat", controllers.GetChatMessagesByCourseID)
        
	}
	return router
}

func AllowCors() gin.HandlerFunc {
    return cors.New(cors.Config{
        AllowOrigins:     []string{"*"}, // Permitir todos los orígenes
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    })
}
