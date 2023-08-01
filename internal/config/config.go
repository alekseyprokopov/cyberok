package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func MustLoad() *Config {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Print("No .env file found")
	}

	return &Config{
		Env:         getEnv("ENV", ""),
		StoragePath: getEnv("STORAGE_PATH", ""),
		HTTPServer: &HTTPServerConfig{
			Address: getEnv("HTTP_SERVER_ADDRESS", ""),
			Timeout: time.Duration(getEnvAsInt("HTTP_SERVER_TIMEOUT", 5)),
		},
		DB: &DBConfig{
			DbUrl: getEnv("DB_URL", ""),
		},
		DNS: &DNSConfig{
			BaseResolvers: getEnvAsSlice("BASE_RESOLVERS", []string{}, ","),
			MaxRetries:    getEnvAsInt("MAX_RETRIES", 5),
			//QuestionTypes: []uint16{miekgdns.TypeA},
		},
	}
}

type Config struct {
	Env         string `env:"ENV"`
	StoragePath string `env:"STORAGE_PATH" env-required:"true"`
	HTTPServer  *HTTPServerConfig
	DB          *DBConfig
	DNS         *DNSConfig
}
type HTTPServerConfig struct {
	Address string        `env:"HTTP_SERVER_ADDRESS" env-required:"true"`
	Timeout time.Duration `env:"HTTP_SERVER_TIMEOUT" env-required:"true"`
}

type DNSConfig struct {
	BaseResolvers []string `env:"BASE_RESOLVERS" env-required:"true"`
	MaxRetries    int      `env:"RETRIES" env-required:"true"`
	//QuestionTypes []uint16 `env:"QUESTION_TYPES" env-required:"true"`
}

type DBConfig struct {
	DbUrl string `env:"DB_URL" env-required:"true"`
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")
	if valStr == "" {
		return defaultVal
	}
	val := strings.Split(valStr, sep)

	return val
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}
