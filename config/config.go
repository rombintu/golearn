package config

import (
	"io/ioutil"
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
	Sqlite3  string
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
}

// Return configuration
func GetConfig(path string) *Config {
	confFile, err := os.Open(path)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer confFile.Close()

	content, err := ioutil.ReadAll(confFile)
	if err != nil {
		log.Fatalf("%v", err)
	}

	var conf Config

	if _, err := toml.Decode(string(content), &conf); err != nil {
		log.Fatalf("%v", err)
	}

	return &conf
}

func GetClientConfig(path string) *ConfigClient {
	confFile, err := os.Open(path)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer confFile.Close()

	content, err := ioutil.ReadAll(confFile)

	if err != nil {
		log.Fatalf("%v", err)
	}

	var conf ConfigClient

	if _, err := toml.Decode(string(content), &conf); err != nil {
		log.Fatalf("%v", err)
	}

	return &conf
}
