package main

import (
	"github.com/kreely/orm-gin/initialisers"
	"github.com/gin-gonic/gin"

)

func init() {
	initialisers.LoadEnvVariables()
	initialisers.ConnectToDB("test.db")
}

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pongo-riffic",
		})
	})

	r.Run()

}
