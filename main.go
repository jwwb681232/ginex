package main

import (
	"ginex/routes"
)

func main() {
	router := routes.Init()
	router.Run()
}
