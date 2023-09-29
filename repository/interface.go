package repository

import (
	"blogspot/entity"
	"database/sql"
)

type IUserRepository interface {
	Login(user entity.User) (entity.User, error)
	Register(user entity.User) (err error)
	Update(user entity.User) (err error)
}

type IBlogRepository interface {
	GetBlog(blogFilter entity.BlogFilter) (blog entity.BlogResponse, err error)
	GetByUser(blogFilter entity.BlogFilter) (blogs []entity.BlogResponse, err error)
	GetHidden(blogFilter entity.BlogFilter) (blogs []entity.BlogResponse, err error)
	Insert(tx *sql.Tx, blog entity.Blog) (lastInsertID int64, err error)
	Update(blog entity.Blog) (err error)
	Hidden(blogFilter entity.BlogFilter) (err error)
	InsertTagDetail(tx *sql.Tx, blogID int64, tagID int) (err error)
	InsertCategoryDetail(tx *sql.Tx, blogID int64, categoryID int) (err error)
	BlogBeginTransaction() (*sql.Tx, error)
	Delete(tx *sql.Tx, blogID int) error
	DeleteTagDetail(tx *sql.Tx, blogID int) (err error)
	DeleteCategoryDetail(tx *sql.Tx, blogID int) (err error)
}

type ITagRepository interface {
	GetAll(tagFilter entity.TagFilter) ([]entity.Tag, error)
	Insert(tag entity.Tag) (err error)
}

type ICategoryRepository interface {
	GetAll(categoryFilter entity.CategoryFilter) ([]entity.Category, error)
	Insert(category entity.Category) (err error)
}
