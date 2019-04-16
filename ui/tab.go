package ui

import (
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

	tabpane := widgets.NewTabPane(dirs...)
	tabpane.Title = "Categories"
	tabpane.SetRect(0, 1, 50, 4)
	tabpane.Border = true

	panes := make([]*Pane, 0, len(directories))
	for _, dir := range directories {
		panes = append(panes, NewPane(dir.Sources()))
	}

	return &Tab{TabPane: tabpane, panes: panes}
}

func (tab *Tab) Refresh() {
	termui.Render(tab.ActivePane())
}

func (tab *Tab) ActivePane() *Pane {
	return tab.panes[tab.TabPane.ActiveTabIndex]
}
