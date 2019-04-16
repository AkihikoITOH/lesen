package ui

import (
	"log"

	termui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"

	"github.com/AkihikoITOH/lesen/model"
)

type TUI struct {
	feed model.Root
}

func NewTUI(root model.Root) *TUI {
	return &TUI{root}
}

func (t *TUI) instruction() *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.Text = "Press q to quit, Press h or l to switch between tabs."
	p.SetRect(0, 0, len(p.Text)+5, 1)
	p.Border = false
	p.TextStyle.Bg = termui.ColorBlue
	return p
}

func (t *TUI) Draw() {
	if err := termui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer termui.Close()

	instruction := t.instruction()
	tab := NewTab(t.feed.Directories())

	termui.Render(instruction, tab.TabPane, tab.panes[0], tab.panes[0].articleLists[0])

	uiEvents := termui.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "h", "<Left>":
			tab.FocusLeft()
			termui.Clear()
			termui.Render(instruction, tab.TabPane)
			tab.Refresh()
			tab.ActivePane().Refresh()
		case "l", "<Right>":
			tab.FocusRight()
			termui.Clear()
			termui.Render(instruction, tab.TabPane)
			tab.Refresh()
			tab.ActivePane().Refresh()
		case "<Up>":
			tab.ActivePane().ScrollUp()
			termui.Clear()
			termui.Render(instruction, tab.TabPane)
			tab.Refresh()
			tab.ActivePane().Refresh()
		case "<Down>":
			tab.ActivePane().ScrollDown()
			termui.Clear()
			termui.Render(instruction, tab.TabPane)
			tab.Refresh()
			tab.ActivePane().Refresh()
		}
	}
}
