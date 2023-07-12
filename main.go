package main

import (
	"fmt"

	"github.com/danielmesquitta/coffee-delivery-api/config"
)

func init() {
	config.InitDotEnv()
}

func main() {
	fmt.Println("Hello, World!")
}
