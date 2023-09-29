package repository

import (
	"blogspot/constant"
	"blogspot/entity"
	"database/sql"
	"strings"

	"github.com/jmoiron/sqlx"
)

type blogRepo struct {
	db *sqlx.DB
}

func NewBlogRepository(db *sqlx.DB) IBlogRepository {
	return &blogRepo{
		db: db,
	}
}

func (r *blogRepo) BlogBeginTransaction() (*sql.Tx, error) {
	return r.db.Begin()
}

func (r *blogRepo) GetBlog(blogFilter entity.BlogFilter) (blog entity.BlogResponse, err error) {
	query := `SELECT bp.id, bp.title, bp.content, u.name FROM users u JOIN blog_posts bp on u.id = bp.user_id WHERE bp.id = ? AND bp.hidden=0`

	err = r.db.Get(&blog, query, blogFilter.ID)
	return
}

func (r *blogRepo) GetByUser(blogFilter entity.BlogFilter) (blogs []entity.BlogResponse, err error) {
	var args []interface{}
	query := `SELECT DISTINCT bp.id, bp.title, bp.content, u.name FROM users u JOIN blog_posts bp on u.id = bp.user_id
			JOIN blog_posts_categories bpc on bp.id = bpc.blog_id
			JOIN blog_posts_tags bpt on bp.id = bpt.blog_id
			JOIN categories c on c.id = bpc.category_id
			JOIN tags t on t.id = bpt.tag_id WHERE u.id = ? AND bp.hidden=0`
	args = append(args, blogFilter.ID)

	if blogFilter.Category != "" || blogFilter.Limit != constant.IntParamEmpty || blogFilter.Offset != constant.IntParamEmpty || blogFilter.Tag != "" {
		query += " AND"
	}

	if blogFilter.Category != "" {
		query += " c.category LIKE (?) AND"
		args = append(args, "%"+blogFilter.Category+"%")
	}

	if blogFilter.Tag != "" {
		query += " t.tag LIKE (?) AND"
		args = append(args, "%"+blogFilter.Tag+"%")
	}

	query = strings.TrimRight(query, "AND")

	if blogFilter.Limit != constant.IntParamEmpty {
		query += " LIMIT ? "
		args = append(args, blogFilter.Limit)
	}

	if blogFilter.Offset != constant.IntParamEmpty {
		query += " OFFSET ?"
		args = append(args, blogFilter.Offset)
	}

	err = r.db.Select(&blogs, query, args...)
	return
}

func (r *blogRepo) GetHidden(blogFilter entity.BlogFilter) (blogs []entity.BlogResponse, err error) {
	var args []interface{}
	query := `SELECT bp.id, bp.title, bp.content, u.name FROM users u JOIN blogspot.blog_posts bp on u.id = bp.user_id WHERE u.id = ? AND bp.hidden=1`
	args = append(args, blogFilter.UserID)

	if blogFilter.Limit != constant.IntParamEmpty {
		query += " LIMIT ? "
		args = append(args, blogFilter.Limit)
	}

	if blogFilter.Offset != constant.IntParamEmpty {
		query += " OFFSET ?"
		args = append(args, blogFilter.Offset)
	}

	err = r.db.Select(&blogs, query, args...)
	return
}

func (r *blogRepo) Insert(tx *sql.Tx, blog entity.Blog) (lastInsertID int64, err error) {
	query := `INSERT INTO blog_posts(title, content, hidden, user_id) VALUES(?,?,?,?)`

	res, err := tx.Exec(query,
		blog.Title,
		blog.Content,
		0,
		blog.UserID,
	)
	if err != nil {
		return
	}

	lastInsertID, err = res.LastInsertId()
	if err != nil {
		return
	}

	return
}

func (r *blogRepo) Update(blog entity.Blog) (err error) {
	query := `UPDATE blog_posts SET title=:title, content=:content WHERE id = :id`

	_, err = r.db.NamedExec(query, blog)

	if err != nil {
		return
	}
	return
}

func (r *blogRepo) Hidden(blogFilter entity.BlogFilter) (err error) {
	query := `UPDATE blog_posts SET hidden=? WHERE id = ?`

	_, err = r.db.Exec(query, blogFilter.Hidden, blogFilter.ID)

	if err != nil {
		return
	}
	return
}

func (r *blogRepo) InsertTagDetail(tx *sql.Tx, blogID int64, tagID int) (err error) {
	query := `INSERT INTO blog_posts_tags(blog_id, tag_id) VALUES (?,?)`

	_, err = tx.Exec(query,
		blogID,
		tagID,
	)

	if err != nil {
		return
	}
	return
}

func (r *blogRepo) InsertCategoryDetail(tx *sql.Tx, blogID int64, categoryID int) (err error) {
	query := `INSERT INTO blog_posts_categories(blog_id, category_id) VALUES (?,?)`

	_, err = tx.Exec(query,
		blogID,
		categoryID,
	)

	if err != nil {
		return
	}
	return
}

func (r *blogRepo) DeleteTagDetail(tx *sql.Tx, blogID int) (err error) {
	query := `DELETE FROM blog_posts_tags WHERE blog_id = ?`

	_, err = tx.Exec(query,
		blogID,
	)

	if err != nil {
		return
	}
	return
}

func (r *blogRepo) DeleteCategoryDetail(tx *sql.Tx, blogID int) (err error) {
	query := `DELETE FROM blog_posts_categories WHERE blog_id = ?`

	_, err = tx.Exec(query,
		blogID,
	)

	if err != nil {
		return
	}
	return
}

func (r *blogRepo) Delete(tx *sql.Tx, blogID int) (err error) {
	query := `DELETE FROM blog_posts WHERE id = ?`

	_, err = tx.Exec(query,
		blogID,
	)

	if err != nil {
		return
	}
	return
}
