package config

import "os"

func LoadConfig() {
	// Load configuration from environment variables or config files
	os.Setenv("DB_USER", "elly")
	os.Setenv("DB_PASSWORD", "dev")
	os.Setenv("DB_NAME", "users")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "3306")

}
