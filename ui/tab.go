package ui

import (
	"fmt"
	"strings"

	"github.com/AkihikoITOH/lesen/model"
	termui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Tab struct {
	*widgets.TabPane
	panes []*Pane
}

func NewTab(directories []model.Directory) *Tab {
	dirs := make([]string, 0, len(directories))
	for _, dir := range directories {
		dirs = append(dirs, dir.Title())
	}

	width := len(fmt.Sprintf(" %s ", strings.Join(dirs, " | ")))

	tabpane := widgets.NewTabPane(dirs...)
	tabpane.SetRect(0, 1, width, 4)
	tabpane.Border = true
	tabpane.BorderStyle.Fg = defaultColor

	panes := make([]*Pane, 0, len(directories))
	for _, dir := range directories {
		panes = append(panes, NewPane(dir.Sources()))
	}

	return &Tab{TabPane: tabpane, panes: panes}
}

func (tab *Tab) Refresh() {
	tab.ActivePane().Refresh()
	termui.Render(tab.ActivePane())
}

func (tab *Tab) ActivePane() *Pane {
	return tab.panes[tab.TabPane.ActiveTabIndex]
}
