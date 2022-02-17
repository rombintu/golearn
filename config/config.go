package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Default struct {
	LogLevel string
	LogFile  string
}

type Server struct {
	Host   string
	Port   string
	Secret string
}

type Postgres struct {
	Dev      bool
	User     string
	Password string
	Dbname   string
	SSLMode  string
}

type Sqlite3 struct {
	Path string
}

type DefaultClient struct {
	Host string `toml:"Host"`
	Port string `toml:"Port"`
}

type Private struct {
	Token string `toml:"Token"`
}
type ConfigClient struct {
	Default DefaultClient
	Private Private
}

type Config struct {
	Default  Default
	Server   Server
	Postgres Postgres
	Sqlite3  Sqlite3
}

// Return configuration
func GetConfig(path string) *Config {
	confFile, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("%v", err)
	}

	var conf Config

	if _, err := toml.Decode(string(confFile), &conf); err != nil {
		log.Fatalf("%v", err)
	}

	return &conf
}

func GetClientConfig(path string) *ConfigClient {
	confFile, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("%v", err)
	}

	var conf ConfigClient

	if _, err := toml.Decode(string(confFile), &conf); err != nil {
		log.Fatalf("%v", err)
	}

	return &conf
}
