package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
	"time"
)

func MustLoad() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Print(err)
		log.Print("No .env file found")
	}

	cfg := Config{}
	if err := envconfig.Process("", &cfg); err != nil {
		log.Print(err)
		log.Print("can't process .env")
	}
	return &cfg
}

type Config struct {
	Env          string `envconfig:"ENV"`
	StoragePath  string `envconfig:"STORAGE_PATH" env-required:"true"`
	MigrationUrl string `envconfig:"MIGRATION_URL" env-required:"true"`
	HTTPServer   *HTTPServerConfig
	DB           *DBConfig
	DNS          *DNSConfig
}
type HTTPServerConfig struct {
	Port        string        `envconfig:"HTTP_SERVER_PORT" env-required:"true"`
	Timeout     time.Duration `envconfig:"HTTP_SERVER_TIMEOUT" env-required:"true"`
	IdleTimeout time.Duration `envconfig:"HTTP_SERVER_IDLE_TIMEOUT" env-required:"true"`
	StopTimeout time.Duration `envconfig:"HTTP_SERVER_STOP_TIMEOUT" env-required:"true"`
}

type DNSConfig struct {
	BaseResolvers []string      `envconfig:"BASE_RESOLVERS" env-required:"true"`
	MaxRetries    int           `envconfig:"RETRIES" env-required:"true"`
	WhoisTimeout  time.Duration `envconfig:"WHOIS_TIMEOUT" env-required:"true"`
}

type DBConfig struct {
	Host        string        `envconfig:"DB_HOST" env-required:"true"`
	Port        string        `envconfig:"DB_PORT" env-required:"true"`
	DBName      string        `envconfig:"DB_NAME" env-required:"true"`
	Username    string        `envconfig:"DB_USER" env-required:"true"`
	Password    string        `envconfig:"DB_PASSWORD" env-required:"true"`
	SSLMode     string        `envconfig:"DB_SSLMODE" env-required:"true"`
	UpdateTimer time.Duration `envconfig:"DB_UPDATE_TIMER" env-required:"true"`
}
