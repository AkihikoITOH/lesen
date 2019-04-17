package ui

import (
	"log"

	termui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/thoas/go-funk"

	"github.com/AkihikoITOH/lesen/model"
)

const (
	defaultColor                 = termui.ColorWhite
	focusedBorderColor           = termui.ColorRed
	selectedTextColor            = termui.ColorRed
	messageBackgroundColor       = termui.ColorBlue
	categoryFocus          focus = iota
	sourceFocus
	articleFocus
)

type focus int

var (
	currentFocus focus = categoryFocus
)

type TUI struct {
	feed        model.Root
	instruction *widgets.Paragraph
	window      *Tab
}

func NewTUI(root model.Root) *TUI {
	instruction := newInstruction()
	tab := NewTab(root.Directories())
	return &TUI{root, instruction, tab}
}

func newInstruction() *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.Text = "Press q to quit, Press h or l to switch between tabs."
	p.SetRect(0, 0, len(p.Text)+5, 1)
	p.Border = false
	p.TextStyle.Bg = messageBackgroundColor
	return p
}

func (t *TUI) initialize() {
	termui.Render(t.instruction, t.window.TabPane, t.window.panes[0], t.window.panes[0].articleLists[0])
}

func (t *TUI) pollEvents() {
	uiEvents := termui.PollEvents()

	for {
		e := <-uiEvents
		if funk.Contains([]string{"c", "s", "a"}, e.ID) {
			t.switchFocus(e.ID)
			t.refresh()
			continue
		}
		switch e.ID {
		case "q", "<C-c>":
			return
		case "h", "<Left>":
			t.window.FocusLeft()
			t.refresh()
		case "l", "<Right>":
			t.window.FocusRight()
			t.refresh()
		case "<Up>":
			t.window.ActivePane().ScrollUp()
			t.refresh()
		case "<Down>":
			t.window.ActivePane().ScrollDown()
			t.refresh()
		}
	}
}

func (t *TUI) switchFocus(keyID string) {
	switch keyID {
	case "c":
		currentFocus = categoryFocus
	case "s":
		currentFocus = sourceFocus
	case "a":
		currentFocus = articleFocus
	}
}

func (t *TUI) refresh() {
	termui.Clear()
	t.window.Refresh()
	termui.Render(t.instruction, t.window.TabPane)
}

func (t *TUI) Draw() {
	if err := termui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer termui.Close()

	t.initialize()
	t.pollEvents()
}
