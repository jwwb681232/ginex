package main

import (
	"encoding/gob"
	"ginex/models/user"
	"ginex/routes"
)

func main() {
	router := routes.Init()
	gob.Register(user.User{})
	router.Static("/resources","./resources")
	_ =router.Run(":8080")
}
