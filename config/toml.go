package config

import (
	_ "embed"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

//go:embed config.toml
var defaultTomlConfig string

const tomlConfigFileName = ".devgo.toml"

func LoadTomlConfig() *Config {
	conf := new(Config)
	dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}

	confPath := filepath.Join(dir, tomlConfigFileName)
	_, err = os.Stat(confPath)
	if os.IsNotExist(err) {
		err = ioutil.WriteFile(confPath, []byte(defaultTomlConfig), 0644)
	}
	if err != nil {
		log.Fatalln(err)
	}

	_, err = toml.DecodeFile(confPath, conf)
	if err != nil {
		log.Fatalln(err)
	}
	return conf
}
