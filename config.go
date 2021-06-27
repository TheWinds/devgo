package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	TabEmojis string        `toml:"tab_emojis"`
	Groups    []GroupConfig `toml:"group"`
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
		log.Fatalln(err)
	}

	confPath := filepath.Join(dir, configFileName)
	_, err = os.Stat(confPath)
	if os.IsNotExist(err) {
		err = ioutil.WriteFile(confPath, []byte(defaultConfig), 0644)
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

const defaultConfig = `# random group tab name preifx emoji
tab_emojis="🐶🐱🐭🦊🐻🐼🐮🐷🐸🐵🦉🦄🐟🐳🐖🐂💥🌈🌞"

[[group]]
# group name
name="tools"
[[group.item]]
# group item title
title="hello"
# group item command to exec
exec="echo hello devgo"
[[group.item]]
title="date now"
exec="date"

[[group]]
name="website"
[[group.item]]
title="github"
exec="open https://github.com"

[[group]]
name="devgo"
[[group.item]]
title="edit config"
exec="vim $HOME/.devgo"`
