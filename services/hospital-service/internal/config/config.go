package config

import (
	"errors"
	"log/slog"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	GRPC           GRPCConfig        `yaml: "grpc"`
	GRPCGateway    GRPCGatewayConfig `yaml: "grpcgateway"`
	Database       DatabaseConfig    `yaml: "database"`
	MigrationsPath string            `yaml: "migrationspath"`
}

type GRPCConfig struct {
	Port string `yaml: "port"`
	Host string `yaml: "host"`
}

type GRPCGatewayConfig struct {
	Port string `yaml: "port"`
	Host string `yaml: "host"`
}

type DatabaseConfig struct {
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func NewConfig(configPath string) (*Config, error) {
	if configPath == "" {
		slog.Error("Config path is empty")
		return nil, errors.New("Config path is empty")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		slog.Error("Config file does not exist: " + configPath)
		return nil, err
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
