package main

import (
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
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
	m.draw()
}

func (m *Menu) hit() {
	Exec(m.items[m.selectedIndex].cmd)
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

var randomEmoij string

func init() {
	rand.Seed(time.Now().Unix())
	var emoijs = []rune("🐶🐱🐭🦊🐻🐼🐮🐷🐸🐵🦉🦄🐟🐳🐖🐂💥🌈🌞")
	n := len(emoijs)
	randomEmoij = string(emoijs[rand.Intn(n)])
}

func (m *Menu) draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	defer termbox.Flush()

	w, h := termbox.Size()

	for i, group := range m.groups {
		color := termbox.ColorDefault
		prefix := " "
		if i == m.selectedGroup {
			prefix = randomEmoij
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

	tbPrint(1, h-2, termbox.ColorWhite, termbox.ColorDefault, "/"+m.input)
	tbPrint(w-10, h-2, termbox.ColorWhite, termbox.ColorDefault, "DevGo")
}

func tbPrint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}
