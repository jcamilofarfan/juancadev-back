package config

import (
	"fmt"
	"os"
)


var (
	PORT = getEnv("PORT", "5000")
	DB = getEnv("DB", "gotodo.db")
	TOKENKEY = getEnv("TOKEN_KEY", "laksdjflkasjfwj92jfslj2qu0-9apsoifjk")
	TOKENEXP = getEnv("TOKEN_EXP", "10h")
)

func getEnv(name string, fallback string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	}

	if fallback != "" {
		return fallback
	}

	panic(fmt.Sprintf(`Environment variable not found :: %v`, name))
}