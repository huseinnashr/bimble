package config

import (
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Version  string   `mapstructure:"version"`
	BaseURL  string   `mapstructure:"base_url"`
	App      App      `mapstructure:"app"`
	Resource Resource `mapstructure:"resource"`
}

type App struct {
	Name    string
	Version string
	HTTP    API `mapstructure:"http"`
	GRPC    API `mapstructure:"grpc"`
}

type API struct {
	Address string        `mapstructure:"address"`
	Timeout time.Duration `mapstructure:"timeout"`
}

type Resource struct {
	SQLDatabase   SQLDatabase   `mapstructure:"sql_database"`
	Redis         Redis         `mapstructure:"redis"`
	OtelCollector OtelCollector `mapstructure:"otel_collector"`
}

type SQLDatabase struct {
	Host     string `mapstructure:"host"`
	Port     int64  `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"db_name"`
}

type Redis struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
}

type OtelCollector struct {
	OTLPGRPC string `mapstructure:"otlp/grpc"`
}

func GetConfig(configPath string) (*Config, error) {
	v := viper.New()
	v.SetConfigType("yaml")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if err := v.ReadConfig(file); err != nil {
		return nil, err
	}

	config := &Config{}
	if err := v.Unmarshal(config); err != nil {
		return nil, err
	}
	return config, nil
}
