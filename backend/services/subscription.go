package services

import (
	"errors"
	"log"

	//"time"
	"backend/clients"
	"backend/dao"

	"gorm.io/gorm"
)

// CreateSubscription crea una nueva suscripción.
func CreateSubscription(subscription dao.Subscription) (dao.Subscription, error) {

    // Busca que la suscricion no exista.
     // Buscar si ya existe una suscripción con IdUsuario e IdCurso
        var existingSubscription dao.Subscription
        result := clients.DB.Where("id_usuario = ? AND id_curso = ?", subscription.IdUsuario, subscription.IdCurso).First(&existingSubscription)

        if result.RowsAffected > 0 {
            // Si ya existe, devolver un mensaje
            return subscription, errors.New(`la suscripción ya existe`)

        }

    tx := clients.DB.Begin()
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()

    // Verificar si el usuario existe
    var user dao.Usuario
    log.Printf("Checking if user exists: %d", subscription.IdUsuario)
    if err := tx.First(&user, subscription.IdUsuario).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            tx.Rollback()
            return subscription, errors.New("user not found")
        }
        tx.Rollback()
        return subscription, err
    }

    // Verificar si el curso existe
    var course dao.Course
    log.Printf("Checking if course exists: %d", subscription.IdCurso)
    if err := tx.First(&course, subscription.IdCurso).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            tx.Rollback()
            return subscription, errors.New("course not found")
        }
        tx.Rollback()
        return subscription, err
    }

    // Crear la suscripción
    //subscription.CreatedAt = time.Now()
    //subscription.UpdatedAt = time.Now()
    if err := tx.Create(&subscription).Error; err != nil {
        tx.Rollback()
        return subscription, err
    }

    if err := tx.Commit().Error; err != nil {
        return subscription, err
    }

    return subscription, nil
}

// GetSubscriptions obtiene todas las suscripciones.
func GetSubscriptions() ([]dao.Subscription, error) {
	var subscriptions []dao.Subscription
	if err := clients.DB.Find(&subscriptions).Error; err != nil {
		return nil, err
	}
	return subscriptions, nil
}

// DeleteSubscription elimina una suscripción por ID.
func DeleteSubscription(id int) error {
	var subscription dao.Subscription
	if err := clients.DB.First(&subscription, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("subscription not found")
		}
		return err
	}
	if err := clients.DB.Delete(&subscription).Error; err != nil {
		return err
	}
	return nil
}
