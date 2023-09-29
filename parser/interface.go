package parser

import (
	"blogspot/entity"
	"net/http"
)

type IUserParser interface {
	ParseUserEntity(r *http.Request) (entity.User, error)
}

type ITagParser interface {
	ParseTagEntity(r *http.Request) (entity.Tag, error)
}

type ICategoryParser interface {
	ParseCategoryEntity(r *http.Request) (entity.Category, error)
}

type IBlogParser interface {
	ParseBlogEntity(r *http.Request) (entity.Blog, error)
	ParseBlogFiler(r *http.Request) entity.BlogFilter
	ParseBlogID(r *http.Request) int
}
