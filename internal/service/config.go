package service

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type BotConfig struct {
	Token        string
	Debug        bool
	Timeout      int
	TemplatePath string
}

type DbConfig struct {
	DriverName string
	Host       string
	Port       int
	Username   string
	Password   string
	Database   string
}

type ApiConfig struct {
	Debug  bool
	AppKey string
}

// Config Application config
type Config struct {
	Version  string
	Bot      BotConfig
	Database struct {
		DriverName string
		Host       string
		Port       int
		Username   string
		Password   string
		Database   string
	}
	Api ApiConfig
}

func ParseConfigFromEnv() *Config {
	c := &Config{}

	c.Version = os.Getenv("APP_VERSION")

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalln("Unable to parse DB port from .env:", err.Error())
	}

	c.Database = DbConfig{
		DriverName: os.Getenv("DB_DRIVER"),
		Host:       os.Getenv("DB_HOST"),
		Port:       dbPort,
		Username:   os.Getenv("DB_USER"),
		Password:   os.Getenv("DB_PASS"),
		Database:   os.Getenv("DB_NAME"),
	}

	if os.Getenv("BOT_TOKEN") == "" {
		log.Panic("BOT_TOKEN is not set")
	}

	c.Bot = BotConfig{
		Token:        os.Getenv("BOT_TOKEN"),
		Debug:        parseBool("BOT_DEBUG", true),
		Timeout:      parseInt("BOT_TIMEOUT", 60),
		TemplatePath: os.Getenv("BOT_TEMPLATE_PATH"),
	}

	c.Api = ApiConfig{
		Debug:  parseBool("API_DEBUG", true),
		AppKey: os.Getenv("API_APPKEY"),
	}

	return c
}

func parseBool(key string, defaultValue bool) bool {
	v, err := strconv.ParseBool(os.Getenv(key))
	if err != nil {
		fmt.Printf("Unable to parse %s from .env: %s\n", key, err.Error())

		v = defaultValue
	}

	return v
}

func parseInt(key string, defaultValue int) int {
	v, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		log.Printf("Unable to parse %s from .env: %s\n", key, err.Error())

		v = defaultValue
	}

	return v
}
