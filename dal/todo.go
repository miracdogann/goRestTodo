package dal

import (
	"fiber_rest/database"

	"gorm.io/gorm"
)

type Todo struct {
	ID        int
	Title     string
	Completed bool `gorm:"default:false"`
}

func CreateTodo(todo *Todo) *gorm.DB {
	return database.DB.Create(&todo)
}

func GetTodos(dest any) *gorm.DB {
	return database.DB.Model(&Todo{}).Find(dest)
}
func GetTodoByID(dest any, id any) *gorm.DB {
	return database.DB.Model(&Todo{}).Where("id = ?", id).First(dest)
}

func UpdateTodo(id any, data any) *gorm.DB {
	return database.DB.Model(&Todo{}).Where("id= ? ", id).Updates(data)
}

func DeleteTodo(id any) *gorm.DB {
	return database.DB.Delete(&Todo{}, id)
}
