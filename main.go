package main

import (
	"fmt"

	"github.com/danielmesquitta/coffee-delivery-api/config"
)

func init() {
	config.LoadEnv()

	config.SetTimeZone()

	config.ConnectToDatabase()

	config.AutoMigrate()
}

func main() {
	fmt.Println("Hello, World!")
}
