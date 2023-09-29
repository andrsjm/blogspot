package parser

import (
	"blogspot/entity"
	"net/http"
)

type categoryParser struct{}

func NewCategoryParser() ICategoryParser {
	return &categoryParser{}
}

func (p *categoryParser) ParseCategoryEntity(r *http.Request) (entity.Category, error) {
	category := entity.Category{}

	err := jsonParser(r, &category)
	if err != nil {
		return category, err
	}

	return category, nil
}
