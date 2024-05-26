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

func GetenvWithDefault(key, defValue string) string {
	str := os.Getenv(key)
	if str != "" {
		return str
	}
	return defValue
}

var SessionSigner = MustGetenvBase64("SESSION_KEY_BASE64")
var DatabaseURL = MustGetenv("DATABASE_URL")
var Port = GetenvWithDefault("PORT", "3000")
var Host = GetenvWithDefault("HOST", "0.0.0.0")
var ListenOn = fmt.Sprintf("%s:%s", Host, Port)

const SessionContextKey = "session"
const FlashContextKey = "flash"
const SessionKey = "_academy_session"
const SessionUserTokenKey = "user_token"
