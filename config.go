package main

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Groups []GroupConfig `toml:"group"`
}

type GroupConfig struct {
	Name  string            `toml:"name"`
	Items []GroupItemConfig `toml:"item"`
}

type GroupItemConfig struct {
	Title string `toml:"title"`
	Exec  string `toml:"exec"`
}

const configFileName = ".devgo"

func LoadConfig() *Config {
	conf := new(Config)
	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	_, err = toml.DecodeFile(filepath.Join(dir, configFileName), conf)
	if err != nil {
		panic(err)
	}
	return conf
}
