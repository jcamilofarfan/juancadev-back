package dal

import (
	"github.com/bancodebogota/bbog-dig-pl-go-mngr-template/src/app/models"
	"github.com/bancodebogota/bbog-dig-pl-go-mngr-template/src/config/database"

	"gorm.io/gorm"
)


func CreateUser(user *models.User) *gorm.DB {
	return database.DB.Create(user)
}

func FindUser(dest interface{}, conds ...interface{}) *gorm.DB {
	return database.DB.Model(&models.User{}).Take(dest, conds...)
}

func FindUserByEmail(dest interface{}, email string) *gorm.DB {
	return FindUser(dest, "email = ?", email)
}
