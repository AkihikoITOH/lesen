package ui

import (
	"github.com/AkihikoITOH/lesen/model"
	termui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type ArticleList struct {
	*widgets.List
}

func NewArticleList(articles []model.Article) *ArticleList {
	arts := make([]string, 0, len(articles))
	for _, art := range articles {
		arts = append(arts, art.Title())
	}

	list := widgets.NewList()
	list.Title = "Articles"
	list.Rows = arts
	list.SetRect(32, 8, 80, 8+2*len(arts))
	list.Border = true
	list.WrapText = false
	list.SelectedRowStyle.Fg = termui.ColorRed

	return &ArticleList{List: list}
}
