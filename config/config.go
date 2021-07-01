package config

import (
	"log"
	"os"
	"os/exec"
)

type Config struct {
	TabEmojis string        `toml:"tab_emojis" yaml:"tab_emojis"`
	Mode      string        `toml:"mode" yaml:"mode"`
	Groups    []GroupConfig `toml:"group" yaml:"group"`
}

type GroupConfig struct {
	Name  string            `toml:"name" yaml:"name"`
	Items []GroupItemConfig `toml:"item" yaml:"item"`
}

type GroupItemConfig struct {
	Title string `toml:"title" yaml:"title"`
	Exec  string `toml:"exec" yaml:"exec"`
}

func LoadConfig() *Config {
	if yaml, ok := FindYaml(); ok {
		return LoadYamlConfig(yaml)
	} else {
		return LoadTomlConfig()
	}
}

func VimConfig() {
	var config string
	if _, ok := FindYaml(); ok {
		config = "$HOME/.devgo.yaml"
	} else {
		config = "$HOME/.devgo.toml"
	}
	cmd := exec.Command("vim", os.ExpandEnv(config))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatalln("vim config Error")
	}
}
