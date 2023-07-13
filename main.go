package main

import (
	"github.com/danielmesquitta/coffee-delivery-api/config"
	"github.com/danielmesquitta/coffee-delivery-api/router"
)

func init() {
	config.Init()
}

func main() {
	router.Init()
}
