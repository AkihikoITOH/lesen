package query

import (
	"github.com/AkihikoITOH/lesen/model"
	"github.com/thoas/go-funk"
)

type CollectDirectoriesOpts struct {
	Titles []string
}

func (q *Query) CollectDirectories(opts CollectDirectoriesOpts) *Query {
	if len(opts.Titles) == 0 {
		return q
	}

	directories := make([]model.Directory, 0)
	for _, dir := range q.Root.Directories() {
		if funk.Contains(opts.Titles, dir.Title()) {
			directories = append(directories, dir)
		}
	}

	newRoot := q.Root.Duplicate()
	newRoot.SetDirectories(directories)

	return &Query{newRoot}
}
