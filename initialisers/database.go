package initialisers

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB(dbFileName string) {
	var err error
	DB, err = gorm.Open(sqlite.Open(dbFileName), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not open DB " + dbFileName)
	}
}
