package services

import (
    "backend/dao"
    "backend/clients"
    "errors"
    "gorm.io/gorm"
)

// GetCourseByID retrieves a course by its ID.
func GetCourseByID(id string) (dao.Course, error) {
    var course dao.Course
    result := clients.DB.Where("id = ?", id).First(&course)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return dao.Course{}, errors.New("record not found")
        }
        return dao.Course{}, result.Error
    }
    return course, nil
}

// GetCoursesByName retrieves courses matching a name pattern.
func GetCoursesByName(name string) ([]dao.Course, error) {
    var courses []dao.Course
    result := clients.DB.Where("nombre LIKE ?", "%"+name+"%").Find(&courses)
    if result.Error != nil {
        return nil, result.Error
    }
    return courses, nil
}

// GetCoursesByCategory retrieves courses by category.

func GetCoursesByCategory(category string) ([]dao.Course, error) {
	var courses []dao.Course
	result := clients.DB.Where("categoria LIKE ?", "%"+category+"%").Find(&courses)
	if result.Error != nil {
		return nil, result.Error
	}
	return courses, nil
}