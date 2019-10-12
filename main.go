package main

import (
	"encoding/gob"
	"ginex/models"
	"ginex/routes"
)

func main() {
	router := routes.Init()
	gob.Register(models.User{})
	router.Static("/resources","./resources")
	_ =router.Run(":8080")
}
