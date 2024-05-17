package config

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

func MustGetenv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		msg := fmt.Sprintf("FATAL: Environment variable %s is not set", key)
		log.Fatal(msg)
	}
	return value
}

func MustGetenvBase64(key string) []byte {
	str := MustGetenv(key)
	value, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		msg := fmt.Sprintf("FATAL: Failed to decode environment variable %s as Base64", key)
		log.Fatal(msg)
	}
	return value
}

var SessionKey = MustGetenvBase64("SESSION_KEY_BASE64")
var DatabaseURL = MustGetenv("DATABASE_URL")
