package repository

import (
	"blogspot/entity"

	"github.com/jmoiron/sqlx"
)

type tagRepo struct {
	db *sqlx.DB
}

func NewTagRepository(db *sqlx.DB) ITagRepository {
	return &tagRepo{
		db: db,
	}
}

func (r *tagRepo) Insert(tag entity.Tag) (err error) {
	query := `INSERT INTO tags(tag) VALUES(:tag)`

	_, err = r.db.NamedExec(query, tag)

	if err != nil {
		return
	}
	return
}

func (r *tagRepo) GetAll(tagFilter entity.TagFilter) (tags []entity.Tag, err error) {
	query := `SELECT t.* FROM tags t JOIN blog_posts_tags bpt on t.id = bpt.tag_id JOIN blog_posts bp on bpt.blog_id = bp.id WHERE bp.id = ?`

	err = r.db.Select(&tags, query, tagFilter.BlogID)

	if err != nil {
		return
	}
	return
}
