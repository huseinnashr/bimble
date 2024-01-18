package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Version  string   `json:"version" yaml:"version"`
	BaseURL  string   `json:"base_url" yaml:"base_url"`
	Server   Server   `json:"serve" yaml:"server"`
	Resource Resource `json:"resource" yaml:"resource"`
}

type Server struct {
	HTTP API `json:"http" yaml:"http"`
	GRPC API `json:"grpc" yaml:"grpc"`
}

type API struct {
	Address string        `mapstructure:"address"`
	Timeout time.Duration `mapstructure:"timeout"`
}

type Resource struct {
	SQLDatabase   SQLDatabase   `json:"sql_database" yaml:"sql_database"`
	Redis         Redis         `json:"redis" yaml:"redis"`
	OtelCollector OtelCollector `json:"otel_collector" yaml:"otel_collector"`
}

type SQLDatabase struct {
	Host     string `json:"host" yaml:"host"`
	Port     int64  `json:"port" yaml:"port"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	DBName   string `json:"db_name" yaml:"db_name"`
}

type Redis struct {
	Address  string `json:"address" yaml:"address"`
	Password string `json:"password" yaml:"password"`
}

type OtelCollector struct {
	OTLPGRPC string `json:"otlp/grpc" yaml:"otlp/grpc"`
}

func GetConfig(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
