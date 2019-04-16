package feedly

import (
	"sync"

	"github.com/AkihikoITOH/lesen/model"
	"github.com/AkihikoITOH/lesen/parser"
)

func NewFeedsFromOPML(path string) (*Root, error) {
	opml, err := parser.LoadOPML(path)
	if err != nil {
		return nil, err
	}

	root := &Root{title: opml.Head.Title}
	for _, outline := range opml.Body.Outlines {
		feed := &Directory{title: outline.Title}
		for _, item := range outline.Outlines {
			source := &Source{title: item.Title, xmlURL: item.XMLURL, htmlURL: item.HTMLURL}
			feed.sources = append(feed.sources, source)
		}
		root.directories = append(root.directories, feed)
	}

	return root, nil
}

func (r *Root) FetchSources() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	sources := make([]model.Source, 0)
	for _, dir := range r.Directories() {
		sources = append(sources, dir.Sources()...)
	}

	for _, src := range sources {
		wg.Add(1)
		go func(s model.Source, g *sync.WaitGroup) {
			s.Fetch()
			g.Done()
		}(src, wg)
	}

	wg.Done()
	wg.Wait()
}

type Root struct {
	title       string
	directories []model.Directory
}

func (r *Root) Duplicate() model.Root {
	dup := *r
	return &dup
}

func (r *Root) Title() string {
	return r.title
}

func (r *Root) SetTitle(title string) {
	r.title = title
}

func (r *Root) Directories() []model.Directory {
	return r.directories
}

func (r *Root) SetDirectories(directories []model.Directory) {
	r.directories = directories
}
