package main

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"github.com/thewinds/devgo/config"
	"log"
	"math/rand"
	"sort"
	"strings"
	"time"
	"unicode"
)

type MenuMode int
type Mode uint8

const (
	EasyMode MenuMode = iota
	VimMode
)

const (
	NormalMode = iota
	FindMode
)

type MenuItem struct {
	title string
	cmd   string
	group string
	score float64
}

type Menu struct {
	Items []*MenuItem

	items         []*MenuItem
	groups        []string
	selectedIndex int
	selectedGroup int
	input         string
	Mode          Mode
}

func (m *Menu) Run(mode MenuMode) {
	switch mode {
	case VimMode:
		m.vimMode()
	case EasyMode:
		m.easyMode()
	}
}

func (m *Menu) easyMode() {
mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEnter:
				termbox.Close()
				err := m.hit()
				if err != nil {
					log.Fatalln(err)
				}
				return
			case termbox.KeyArrowUp:
				m.selectPrev()
			case termbox.KeyArrowDown:
				m.selectNext()
			case termbox.KeyArrowLeft:
				m.selectPreGroup()
			case termbox.KeyArrowRight, termbox.KeyTab:
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
			termbox.Close()
			log.Fatalln(ev.Err)

		case termbox.EventInterrupt:
			break mainloop
		}
	}
}

func (m *Menu) vimMode() {
mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEnter:
				termbox.Close()
				err := m.hit()
				if err != nil {
					log.Fatalln(err)
				}
				return
			case termbox.KeyArrowUp:
				m.selectPrev()
			case termbox.KeyArrowDown:
				m.selectNext()
			case termbox.KeyArrowLeft:
				m.selectPreGroup()
			case termbox.KeyArrowRight, termbox.KeyTab:
				m.selectNextGroup()
			case termbox.KeyEsc, termbox.KeyCtrlC, termbox.KeyCtrlD:
				break mainloop
			case termbox.KeyBackspace, termbox.KeyBackspace2, termbox.KeyDelete:
				m.filter("")
			default:
				switch ev.Ch {
				case 'j':
					m.selectNext()
				case 'k':
					m.selectPrev()
				case 'h':
					m.selectPreGroup()
				case 'l':
					m.selectNextGroup()
				case 'i':
					termbox.Close()
					config.VimConfig()
					err := termbox.Init()
					if err != nil {
						log.Fatalln(err)
					}
					m.Update(config.LoadConfig())
				case '/':
					m.Mode = FindMode
					m.draw()
				findLoop:
					for {
						switch ev := termbox.PollEvent(); ev.Key {
						case termbox.KeyEsc:
							break findLoop
						case termbox.KeyEnter:
							termbox.Close()
							err := m.hit()
							if err != nil {
								log.Fatalln(err)
							}
							return
						case termbox.KeyBackspace, termbox.KeyBackspace2, termbox.KeyDelete:
							m.filter("")
						default:
							if ev.Type == 0 {
								m.filter(string(ev.Ch))
							}
						}
					}
					m.Mode = NormalMode
					m.draw()
				}

			}

		case termbox.EventError:
			termbox.Close()
			log.Fatalln(ev.Err)

		case termbox.EventInterrupt:
			break mainloop
		}
	}
}

func (m *Menu) init() {
	mark := make(map[string]bool)
	m.groups = nil
	for _, item := range m.Items {
		exist := mark[item.group]
		if !exist {
			m.groups = append(m.groups, item.group)
			mark[item.group] = true
		}
	}
	m.selectGroup(m.selectedGroup)
}

func (m *Menu) selectNext() {
	m.selectedIndex = (m.selectedIndex + 1) % len(m.items)
	m.draw()
}

func (m *Menu) selectPrev() {
	m.selectedIndex = (m.selectedIndex - 1 + len(m.items)) % len(m.items)
	m.draw()
}

func (m *Menu) selectNextGroup() {
	m.selectedGroup = (m.selectedGroup + 1) % len(m.groups)
	m.selectGroup(m.selectedGroup)
}

func (m *Menu) selectPreGroup() {
	m.input = ""
	m.selectedGroup = (m.selectedGroup - 1 + len(m.groups)) % len(m.groups)
	m.selectGroup(m.selectedGroup)
}

func (m *Menu) selectGroup(i int) {
	name := m.groups[i]
	var items []*MenuItem
	for _, item := range m.Items {
		if item.group == name {
			items = append(items, item)
		}
	}
	m.items = items
	m.selectedIndex = 0
	m.draw()
}

func (m *Menu) hit() error {
	return Exec(m.items[m.selectedIndex].cmd)
}

func (m *Menu) filter(s string) {
	if len(s) == 0 {
		if len(m.input) != 0 {
			m.input = m.input[:len(m.input)-1]
		}
	} else {
		m.input += s
	}
	m.init()
	if len(m.input) == 0 {
		m.draw()
		return
	}
	for _, item := range m.Items {
		score := float64(strings.Count(item.title, m.input))
		for _, term := range m.input {
			score += float64(strings.Count(item.title, string(term))) / float64(len(item.title))
		}
		item.score = score
	}
	sort.Slice(m.items, func(i, j int) bool {
		return m.items[i].score > m.items[j].score
	})
	m.draw()
}

var randomEmoji string

func initTabEmojis(conf *config.Config) {
	if len(conf.TabEmojis) == 0 {
		return
	}
	rand.Seed(time.Now().Unix())
	var emojisRune = []rune(conf.TabEmojis)
	n := len(emojisRune)
	randomEmoji = string(emojisRune[rand.Intn(n)])
}

func (m *Menu) draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	defer termbox.Flush()

	w, h := termbox.Size()

	for i, group := range m.groups {
		color := termbox.ColorDefault
		prefix := " "
		if i == m.selectedGroup {
			prefix = randomEmoji
			color = termbox.ColorLightGreen
		}
		tbPrint(2+i*10, 1, color, termbox.ColorDefault, prefix+group)
	}

	for i, item := range m.items {
		selected := m.selectedIndex == i
		prefix := "  "
		color := termbox.ColorDefault
		if selected {
			prefix = "> "
			color = termbox.ColorLightGreen
		}
		tbPrint(1, i+3, color, termbox.ColorDefault, prefix+item.title)
	}
	if m.Mode == FindMode {
		tbPrint(1, h-2, termbox.ColorWhite, termbox.ColorDefault, "/"+m.input)
		tbPrint(w-10, h-2, termbox.ColorWhite, termbox.ColorDefault, "DevGo")
	}
}

func (m *Menu) Update(conf *config.Config) {
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
	m.Items = items
	m.init()
	m.draw()
}

func tbPrint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}
