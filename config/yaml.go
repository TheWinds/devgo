package config

import (
	_ "embed"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"path/filepath"
)

//go:embed config.yaml
var defaultYaml string

const yamlConfigFileName = ".devgo.yaml"

func FindYaml() (string, bool) {
	dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}
	confPath := filepath.Join(dir, yamlConfigFileName)
	_, err = os.Stat(confPath)
	if os.IsNotExist(err) {
		return confPath, false
	}
	if err != nil {
		log.Fatalln(err)
	}
	return confPath, true
}

func LoadYamlConfig(configPath string) *Config {
	conf := new(Config)

	fileContent, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalln(err)
	}

	err = yaml.Unmarshal(fileContent, conf)
	if err != nil {
		log.Fatalln(err)
	}
	return conf
}
