package ui

import (
	"github.com/AkihikoITOH/lesen/model"
	termui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Pane struct {
	*widgets.List
	articleLists []*ArticleList
}

func NewPane(sources []model.Source) *Pane {
	srcs := make([]string, 0, len(sources))
	for _, src := range sources {
		srcs = append(srcs, src.Title())
	}

	list := widgets.NewList()
	list.Title = "Sources"
	list.Rows = srcs
	list.SetRect(0, 8, 30, 8+2*len(srcs))
	list.Border = true
	list.WrapText = false
	list.SelectedRowStyle.Fg = termui.ColorRed

	articleLists := make([]*ArticleList, 0, len(sources))
	for _, src := range sources {
		articleLists = append(articleLists, NewArticleList(src.Articles()))
	}

	return &Pane{List: list, articleLists: articleLists}
}

func (pane *Pane) Refresh() {
	termui.Render(pane.ActiveArticleList())
}

func (pane *Pane) ActiveArticleList() *ArticleList {
	return pane.articleLists[pane.List.SelectedRow]
}
