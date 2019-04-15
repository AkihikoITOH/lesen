package ui

import (
	"fmt"

	"github.com/AkihikoITOH/lesen/model"
)

const (
	DirectoryLevel NestLevel = iota
	SourceLevel
	ArticleLevel
)

type NestLevel int

type CUI struct {
	feed  model.Root
	level NestLevel
}

func NewCUI(root model.Root, level NestLevel) *CUI {
	return &CUI{root, level}
}

func (c *CUI) Draw() {
	c.drawDirectories(c.feed.Directories())
}

func (c *CUI) drawDirectories(directories []model.Directory) {
	for _, dir := range c.feed.Directories() {
		fmt.Printf("* %s\n", dir.Title())
		c.drawSources(dir.Sources())
	}
}

func (c *CUI) drawSources(sources []model.Source) {
	if c.level < SourceLevel {
		return
	}
	for _, src := range sources {
		fmt.Printf(" * %s\n", src.Title())
		c.drawArticles(src.Articles())
	}
}

func (c *CUI) drawArticles(articles []model.Article) {
	if c.level < ArticleLevel {
		return
	}
	for _, article := range articles {
		if len(article.Title()) > 0 {
			fmt.Printf("  * %s\n", article.Title())
			fmt.Printf("    %s\n", article.Link())
		}
	}
}
