package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

type Config struct {
	Host    string        `json:"host"`
	Port    int           `json:"port"`
	Debug   bool          `json:"debug"`
	Cache   string        `json:"cache"`
	LevelDB LevelDBConfig `json:"leveldb"`
	Env     string
}

type LevelDBConfig struct {
	Path string `json:"path"`
}

var config *Config

func GetEnv() string {
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	return strings.ToLower(env)
}

func GetConfig() *Config {
	if config == nil {
		env := GetEnv()

		config = &Config{
			Host:    "127.0.0.1",
			Port:    8080,
			Debug:   env != "production",
			Cache:   "leveldb",
			LevelDB: LevelDBConfig{Path: "cache.db"},
			Env:     env,
		}

		fileName := "config/" + env + ".json"
		if data, err := ioutil.ReadFile(fileName); err == nil {
			json.Unmarshal(data, &config)
		}
	}

	return config
}
