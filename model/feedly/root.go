package feedly

import (
	"sync"

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
	for _, dir := range r.Directories() {
		for _, src := range dir.Sources() {
			wg.Add(1)
			go func(s *Source, g *sync.WaitGroup) {
				s.Fetch()
				g.Done()
			}(src, wg)
		}
	}
	wg.Done()
	wg.Wait()
}

type Root struct {
	title       string
	directories []*Directory
}

func (r *Root) Title() string {
	return r.title
}

func (r *Root) Directories() []*Directory {
	return r.directories
}
