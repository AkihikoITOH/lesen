package feedly

import (
	"sync"

	"gopkg.in/cheggaaa/pb.v1"

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

	progress := pb.StartNew(len(sources))
	for _, src := range sources {
		wg.Add(1)
		go func(s model.Source, g *sync.WaitGroup, p *pb.ProgressBar) {
			s.Fetch()
			p.Increment()
			g.Done()
		}(src, wg, progress)
	}

	wg.Done()
	wg.Wait()
}

type Root struct {
	title       string
	directories []model.Directory
}

func (r *Root) Title() string {
	return r.title
}

func (r *Root) Directories() []model.Directory {
	return r.directories
}
