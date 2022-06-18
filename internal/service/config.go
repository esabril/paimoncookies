package service

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
	"log"
)

const EnvVariablesPrefix = "PCOOKIES_"

type BotConfig struct {
	Token        string `env:"BOT_TOKEN,required"`
	Debug        bool   `env:"BOT_DEBUG"`
	Timeout      int    `env:"BOT_TIMEOUT,default=60"`
	TemplatePath string `env:"BOT_TEMPLATE_PATH,required"`
}

type DbConfig struct {
	DriverName string `env:"DB_DRIVER,required"`
	Host       string `env:"DB_HOST,required"`
	Port       int    `env:"DB_PORT,required"`
	Username   string `env:"DB_USER,required"`
	Password   string `env:"DB_PASS,required"`
	Database   string `env:"DB_NAME,required"`
	SslMode    string `env:"DB_SSL_MODE,required"`
}

type ApiConfig struct {
	Debug  bool   `env:"API_DEBUG"`
	AppKey string `env:"API_APPKEY,required"`
	Port   string `env:"API_PORT,required"`
}

// Config Application config
type Config struct {
	Version  string `env:"APP_VERSION,required"`
	Timezone string `env:"TIMEZONE,default=Asia/Almaty"`
	Bot      BotConfig
	Database DbConfig
	Api      ApiConfig
}

func ParseConfigFromEnv(ctx context.Context) *Config {
	c := &Config{}

	// For convenience, you can use the .env file in the root directory of the project
	err := godotenv.Load()
	if err != nil {
		log.Printf("%s. Will use config from environment\n", err.Error())
	}

	l := envconfig.PrefixLookuper(EnvVariablesPrefix, envconfig.OsLookuper())
	if err := envconfig.ProcessWith(ctx, c, l); err != nil {
		panic(err)
	}

	return c
}
