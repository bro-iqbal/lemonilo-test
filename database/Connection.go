package database

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connection() *gorm.DB {
	db, err := gorm.Open("mysql", os.Getenv("MYSQL_DB_USERNAME")+":"+os.Getenv("MYSQL_DB_PASSWORD")+"@tcp("+os.Getenv("MYSQL_DB_HOST")+":"+os.Getenv("MYSQL_DB_PORT")+")/"+os.Getenv("MYSQL_DB_DATABASE")+"?charset=utf8&parseTime=True")
	if err != nil {
		log.Println("[*] Open DB connection successful")
	}

	log.Println("[*] Open DB connection successful")

	return db
}
