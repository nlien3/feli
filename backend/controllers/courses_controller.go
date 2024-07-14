package controllers

import (
	"backend/clients"
	"backend/dao"
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCourses(c *gin.Context) {
	courses, err := services.GetCourses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, courses)
}

func CreateCourse(c *gin.Context) {
	var course dao.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newCourse, err := services.CreateCourse(course)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newCourse)
}


func DeleteCourse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	// Verify that the course exists
	var course dao.Course
	if err := clients.DB.First(&course, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "The specified course does not exist"})
		return
	}

	// Delete the subscriptions associated with the course
	if err := clients.DB.Where("Id_curso = ?", id).Delete(&dao.Subscription{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete subscriptions"})
		return
	}

	// Delete the course
	if err := clients.DB.Delete(&course).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the course"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course and associated subscriptions successfully deleted"})
}


func GetSuscriptionByIdUser(c *gin.Context) {
	id := c.Param("id")
	course, err := services.GetCourseByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}
	c.JSON(http.StatusOK, course)
}


func UpdateCourseHandler(c *gin.Context) {
    var course dao.Course
    id := c.Param("id")

    if err := clients.DB.First(&course, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
        return
    }

    if err := c.ShouldBindJSON(&course); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := clients.DB.Save(&course).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update course"})
        return
    }

    c.JSON(http.StatusOK, course)
}
