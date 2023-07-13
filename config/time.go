package config

import (
	"log"
	"os"
	"time"
)

func setTimeZone() {
	var err error

	timezone := os.Getenv("TIMEZONE")

	if timezone == "" {
		log.Fatal("timezone not set in .env file")
	}

	time.Local, err = time.LoadLocation(timezone)

	if err != nil {
		log.Fatal("failed to set timezone", err)
	}
}
