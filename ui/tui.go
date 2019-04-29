package ui

import (
	"log"

	tm "github.com/buger/goterm"
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
	sourceFocus            focus = iota
	articleFocus
)

type focus int

var (
	currentFocus focus = sourceFocus
)

type TUI struct {
	feed                model.Root
	globalInstruction   *widgets.Paragraph
	window              *Tab
	detailedInstruction *widgets.Paragraph
}

func NewTUI(root model.Root) *TUI {
	tm.Clear()
	globalInstruction := newGlobalInstruction()
	detailedInstruction := newDetailedInstruction()
	tab := NewTab(root.Directories())
	return &TUI{root, globalInstruction, tab, detailedInstruction}
}

func newGlobalInstruction() *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.Text = "Switch tabs: \"Left\", \"Right\" | Exit: \"q\""
	p.SetRect(0, 0, len(p.Text)+5, 1)
	p.Border = false
	p.TextStyle.Bg = messageBackgroundColor
	return p
}

func newDetailedInstruction() *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.Text = "Switch focus: \"s\" (sources), \"a\" (articles) | Browse items: \"Up\", \"Down\" | Open in browser: \"Enter\""
	p.SetRect(0, 5, len(p.Text)+5, 6)
	p.Border = false
	p.TextStyle.Bg = messageBackgroundColor
	return p
}

func (t *TUI) initialize() {
	termui.Render(t.globalInstruction, t.window.TabPane, t.detailedInstruction, t.window.panes[0], t.window.panes[0].articleLists[0])
}

func (t *TUI) pollEvents() {
	uiEvents := termui.PollEvents()

	for {
		e := <-uiEvents
		// Quit application
		if funk.Contains([]string{"q", "<C-c>"}, e.ID) {
			break
		}
		// Switch focus
		if funk.Contains([]string{"c", "s", "a"}, e.ID) {
			t.switchFocus(e.ID)
			continue
		}
		// Switch tab
		if funk.Contains([]string{"<Left>", "<Right>"}, e.ID) {
			t.switchTab(e.ID)
			continue
		}
		// Browse
		switch currentFocus {
		case sourceFocus:
			t.browseSources(e.ID)
		case articleFocus:
			t.browseArticles(e.ID)
		}
	}
}

func (t *TUI) switchTab(eventID string) {
	defer t.refresh()

	switch eventID {
	case "<Left>":
		t.window.FocusLeft()
	case "<Right>":
		t.window.FocusRight()
	}
}

func (t *TUI) browseSources(eventID string) {
	defer t.refresh()

	switch eventID {
	case "<Up>":
		t.window.ActivePane().ScrollUp()
	case "<Down>":
		t.window.ActivePane().ScrollDown()
	}
}

func (t *TUI) browseArticles(eventID string) {
	defer t.refresh()

	switch eventID {
	case "<Up>":
		t.window.ActivePane().ActiveArticleList().ScrollUp()
	case "<Down>":
		t.window.ActivePane().ActiveArticleList().ScrollDown()
	}
}

func (t *TUI) switchFocus(keyID string) {
	defer t.refresh()

	switch keyID {
	case "s":
		currentFocus = sourceFocus
	case "a":
		currentFocus = articleFocus
	}
}

func (t *TUI) refresh() {
	tm.Clear()
	termui.Clear()
	t.window.Refresh()
	termui.Render(t.globalInstruction, t.window.TabPane, t.detailedInstruction)
}

func (t *TUI) Draw() {
	if err := termui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer termui.Close()

	t.initialize()
	t.pollEvents()
}
