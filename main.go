package main

import (
	"ginex/routes"
)

func main() {
	router := routes.Init()
	router.Static("/resources","./resources")
	router.Run()
}
