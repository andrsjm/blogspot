package parser

import (
	"blogspot/entity"
	"net/http"
)

type tagParser struct{}

func NewTagParser() ITagParser {
	return &tagParser{}
}

func (p *tagParser) ParseTagEntity(r *http.Request) (entity.Tag, error) {
	tag := entity.Tag{}

	err := jsonParser(r, &tag)
	if err != nil {
		return tag, err
	}

	return tag, nil
}
