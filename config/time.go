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
		log.Fatal("Timezone not set in .env file")
	}

	time.Local, err = time.LoadLocation(timezone)

	if err != nil {
		log.Fatal("Failed to set timezone", err)
	}
}
