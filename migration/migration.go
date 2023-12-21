package migration

import (
	"github.com/abhis3110/go-todoApp/model"
	//"github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB){
	db.AutoMigrate(&model.Task{})
	db.AutoMigrate(&model.User{})
}