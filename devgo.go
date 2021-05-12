package main

import (
	"unicode"

	"github.com/nsf/termbox-go"
)

func main() {
	conf := LoadConfig()
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
	m := Menu{Items: items}

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	termbox.SetInputMode(termbox.InputEsc)
	m.init()
	m.draw()

mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEnter:
				termbox.Close()
				m.hit()
				return
			case termbox.KeyArrowUp:
				m.selectPrev()
			case termbox.KeyArrowDown:
				m.selectNext()
			case termbox.KeyArrowLeft:
				m.selectPreGroup()
			case termbox.KeyArrowRight:
				m.selectNextGroup()
			case termbox.KeyEsc, termbox.KeyCtrlC, termbox.KeyCtrlD:
				break mainloop
			case termbox.KeyBackspace, termbox.KeyBackspace2, termbox.KeyDelete:
				m.filter("")
			default:
				if unicode.IsLetter(ev.Ch) {
					m.filter(string(ev.Ch))
				}
			}

		case termbox.EventError:
			panic(ev.Err)

		case termbox.EventInterrupt:
			break mainloop
		}
	}
	termbox.Close()
}
