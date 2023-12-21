package config

import (
	//"github.com/jinzhu/gorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const ( 
	IdentityKey = "id"
	Key = "my_secret_key_8F6E2P"
)


var DB *gorm.DB

func CreatePostgraceDB() *gorm.DB{
	dsn:="host=postgrestodo port=5432 user=admin password=123 dbname=tododb sslmode=disable"
	db, err:=gorm.Open(postgres.Open(dsn), &gorm.Config{})

	 if err!=nil{
		panic(err.Error())
	 }

	 DB=db
	 return db
}


 // utility function to get database reference 
func GetDB() *gorm.DB {
	return DB
}