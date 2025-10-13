package initialisers

import (
	"gorm.io/drivers/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB(dbFileName string) {
	var err error
	filename := "recipes.db"
	db, err := gorm.Open(sqlite.Open(dbFileName), &gorm.Config{})
}
