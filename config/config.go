package config

import (
	"github.com/thewinds/devgo/utils"
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
	UpdateOldConfig()
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

//UpdateOldConfig update .devgo to .devgo.toml
// We will delete this function in the future
func UpdateOldConfig() {
	oldPath := os.ExpandEnv(`$HOME/.devgo`)
	newPath := os.ExpandEnv(`$HOME/.devgo.toml`)
	if !utils.Exists(oldPath) {
		return
	}

	err := os.Rename(oldPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}
