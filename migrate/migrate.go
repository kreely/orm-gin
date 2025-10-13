package main

import (
	"github.com/kreely/org-gin/initialisers"
	"github.com/kreely/org-gin/models"
)

func init() {
	initialisers.loadEnvVariables() 
	initialisers.ConnectToDB()
}

func main() {
	initialisers.DB.AutoMigrate(&models.Post{})
}
