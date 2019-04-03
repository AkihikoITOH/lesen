package feedly

import (
	"github.com/AkihikoITOH/lesen/model"
)

type Directory struct {
	title   string
	sources []model.Source
}

func (f *Directory) Title() string {
	return f.title
}

func (f *Directory) Sources() []model.Source {
	return f.sources
}
