package flow

import "blogspot/entity"

type IUserFlow interface {
	Login(user entity.User) (entity.User, error)
	Register(user entity.User) (err error)
	Update(user entity.User) (err error)
}

type ITagFlow interface {
	Insert(tag entity.Tag) (err error)
}

type ICategoryFlow interface {
	Insert(category entity.Category) (err error)
}

type IBlogFlow interface {
	GetBlog(blogFilter entity.BlogFilter) (blog entity.BlogResponse, err error)
	GetByUser(blogFilter entity.BlogFilter) (blogs []entity.BlogResponse, err error)
	GetHidden(blogFilter entity.BlogFilter) (blogs []entity.BlogResponse, err error)
	Insert(blog entity.Blog) (err error)
	Update(blog entity.Blog) (err error)
	Hidden(blogFilter entity.BlogFilter) (err error)
	Delete(blogID int) (err error)
}
