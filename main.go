package main

import (
	"github.com/danielmesquitta/coffee-delivery-api/config"
	"github.com/danielmesquitta/coffee-delivery-api/controller"
	"github.com/danielmesquitta/coffee-delivery-api/router"
)

func init() {
	config.Init()
	controller.Init()
}

func main() {
	router.Init()
}
