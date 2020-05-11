package store

import (
	"github.com/jinzhu/gorm"
	"log"
)

func GetDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=hoorzy dbname=hoorzy password=sahar67 sslmode=disable")
	if err != nil {
		log.Println("shred.store package.  GetDB error is: ", err)
		return nil

	}
	return db
}
