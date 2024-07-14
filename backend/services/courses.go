package services

import (
	"backend/clients"
	"backend/dao"
)

func GetCourses() ([]dao.Course, error) {
	var courses []dao.Course
	result := clients.DB.Find(&courses)
	return courses, result.Error
}

func CreateCourse(course dao.Course) (dao.Course, error) {
	result := clients.DB.Create(&course)
	return course, result.Error
}

func DeleteCourse(id int) error {
	result := clients.DB.Delete(&dao.Course{}, id)
	return result.Error
}
