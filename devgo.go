package main

import (
	"github.com/nsf/termbox-go"
	"github.com/thewinds/devgo/config"
	"log"
)

func main() {
	conf := config.LoadConfig()
	initTabEmojis(conf)
	m := initMenu(conf)

	err := termbox.Init()
	if err != nil {
		log.Fatalln(err)
	}
	termbox.SetInputMode(termbox.InputEsc)

	m.init()
	m.draw()
	m.Run(conf.Mode)
	termbox.Close()
}

func initMenu(conf *config.Config) *Menu {
	var items []*MenuItem
	for _, group := range conf.Groups {
		for _, item := range group.Items {
			items = append(items, &MenuItem{
				title: item.Title,
				cmd:   item.Exec,
				group: group.Name,
			})
		}
	}
	return &Menu{Items: items}
}
