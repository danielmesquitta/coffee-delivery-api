package config

func Init() {
	loadEnv()
	setTimeZone()
	connectToDatabase()
	autoMigrate()
}
