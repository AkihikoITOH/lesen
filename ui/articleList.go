package ui

import (
	"github.com/AkihikoITOH/lesen/model"
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
	list.SetRect(32, 5, 180, 8+2*len(arts))
	list.Border = true
	list.WrapText = false
	list.SelectedRowStyle.Fg = selectedTextColor

	return &ArticleList{List: list}
}

func (al *ArticleList) Refresh() {
	if currentFocus == articleFocus {
		al.List.BorderStyle.Fg = focusedBorderColor
	} else {
		al.List.BorderStyle.Fg = defaultColor
	}
}
