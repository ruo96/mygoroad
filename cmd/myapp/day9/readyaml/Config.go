package Config

import "time"

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type DatabaseConfig struct {
	User     string          `yaml:"user"`
	Password string          `yaml:"password"`
	Host     string          `yaml:"host"`
	Port     int             `yaml:"port"`
	Name     string          `yaml:"name"`
	Options  DatabaseOptions `yaml:"options"`
}

type DatabaseOptions struct {
	Timeout time.Duration `yaml:"timeout"`
}
