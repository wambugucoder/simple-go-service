package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/wambugucoder/simple-go-service/configs"
	"github.com/wambugucoder/simple-go-service/models"
	"log"
)

func DoesEmailExist(email string) bool {
	var user models.User
	log.Println("addada")
	err := configs.DB.Where(&models.User{Email: email}).First(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false
		}
	}
	return true
}

func SaveUser(user *models.User) bool {
	err := configs.DB.Create(user).Error
	if err != nil {
		return false
	}
	return true

}
