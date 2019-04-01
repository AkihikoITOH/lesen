package feedly

type Directory struct {
	title   string
	sources []*Source
}

func (f *Directory) Title() string {
	return f.title
}

func (f *Directory) Sources() []*Source {
	return f.sources
}
