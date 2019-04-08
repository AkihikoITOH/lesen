package parser

import (
	"github.com/AkihikoITOH/lesen/config"
	"github.com/gilliek/go-opml/opml"
)

func LoadOPML(path string) (*opml.OPML, error) {
	doc, err := opml.NewOPMLFromFile(path)
	if err != nil {
		config.Logger().Warnf(err.Error())
		return doc, err
	}
	return doc, nil
}
