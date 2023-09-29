package parser

import (
	"blogspot/entity"
	"net/http"
)

type blogParser struct{}

func NewBlogParser() IBlogParser {
	return &blogParser{}
}

func (p *blogParser) ParseBlogEntity(r *http.Request) (entity.Blog, error) {
	blog := entity.Blog{}

	err := jsonParser(r, &blog)
	if err != nil {
		return blog, err
	}

	blog.UserID = getUserID(r)

	return blog, nil
}

func (p *blogParser) ParseBlogFiler(r *http.Request) entity.BlogFilter {
	req := clientRequest{r}

	filter := entity.BlogFilter{}
	filter.Offset = req.paramInt("offset")
	filter.Limit = req.paramInt("limit")
	filter.Category = req.param("category")
	filter.Tag = req.param("tag")
	filter.Hidden = req.paramInt("hidden")
	filter.UserID = getUserID(r)

	filter.ID = p.ParseBlogID(r)

	return filter
}

func (p *blogParser) ParseBlogID(r *http.Request) int {
	req := clientRequest{r}
	return req.getEntityID()
}
