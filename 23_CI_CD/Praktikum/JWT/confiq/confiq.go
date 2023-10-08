package confiq

import (

	"fmt"
	
	"mwr/model"
	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
	InitialMigration()
}


func InitDB() {
	config := model.Db{
		DB_Username: "root",
		DB_Password: "",
		DB_Port:     "3306",
		DB_Host:     "35.193.174.90",
		DB_Name:     "crud_go",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)


	var err error
	DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Book{})
	DB.AutoMigrate(&model.Blog{})
}