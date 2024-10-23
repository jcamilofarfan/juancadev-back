package dal

import (
	"github.com/bancodebogota/bbog-dig-pl-go-mngr-template/src/app/models"
	"github.com/bancodebogota/bbog-dig-pl-go-mngr-template/src/config/database"


	"gorm.io/gorm"
)

func CreateTodo(todo *models.Todo) *gorm.DB {
	return database.DB.Create(todo)
}

func FindTodo(dest interface{}, conds ...interface{}) *gorm.DB {
	return database.DB.Model(&models.Todo{}).Take(dest, conds...)
}

func FindTodoByUser(dest interface{}, todoIden interface{}, userIden interface{}) *gorm.DB {
	return FindTodo(dest, "id = ? AND user = ?", todoIden, userIden)
}

func FindTodosByUser(dest interface{}, userIden interface{}) *gorm.DB {
	return database.DB.Model(&models.Todo{}).Find(dest, "user = ?", userIden)
}

func DeleteTodo(todoIden interface{}, userIden interface{}) *gorm.DB {
	return database.DB.Unscoped().Delete(&models.Todo{}, "id = ? AND user = ?", todoIden, userIden)
}

func UpdateTodo(todoIden interface{}, userIden interface{}, data interface{}) *gorm.DB {
	return database.DB.Model(&models.Todo{}).Where("id = ? AND user = ?", todoIden, userIden).Updates(data)
}
