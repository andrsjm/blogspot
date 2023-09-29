package repository

import (
	"blogspot/entity"

	"github.com/jmoiron/sqlx"
)

type categoryRepo struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) ICategoryRepository {
	return &categoryRepo{
		db: db,
	}
}

func (r *categoryRepo) Insert(category entity.Category) (err error) {
	query := `INSERT INTO categories(category) VALUES(:category)`

	_, err = r.db.NamedExec(query, category)

	if err != nil {
		return
	}
	return
}

func (r *categoryRepo) GetAll(categoryFilter entity.CategoryFilter) (categories []entity.Category, err error) {
	query := `SELECT c.* FROM categories c JOIN blog_posts_categories bpc on c.id = bpc.category_id JOIN blog_posts bp on bp.id = bpc.blog_id WHERE bp.id=?`

	err = r.db.Select(&categories, query, categoryFilter.BlogID)

	if err != nil {
		return
	}
	return
}
