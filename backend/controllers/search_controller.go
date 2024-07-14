package controllers

import (
	"backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SearchByID(c *gin.Context) {
	id := c.Param("id")
	course, err := services.GetCourseByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}
	c.JSON(http.StatusOK, course)
}

func SearchByName(c *gin.Context) {
	name := c.Param("name")
	courses, err := services.GetCoursesByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching courses"})
		return
	}
	if len(courses) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No courses found"})
		return
	}
	c.JSON(http.StatusOK, courses)
}

func GetCoursesByCategory(c *gin.Context) {
	category := c.Param("category")
	courses, err := services.GetCoursesByCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching courses by category"})
		return
	}
	if len(courses) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No courses found in this category"})
		return
	}
	c.JSON(http.StatusOK, courses)
}