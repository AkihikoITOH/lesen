package parser

import (
	"github.com/gilliek/go-opml/opml"
	"github.com/sirupsen/logrus"
)

func LoadOPML(path string) (*opml.OPML, error) {
	doc, err := opml.NewOPMLFromFile(path)
	if err != nil {
		logrus.Warnf(err.Error())
		return doc, err
	}
	return doc, nil
}
