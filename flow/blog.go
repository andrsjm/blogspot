package flow

import (
	"blogspot/entity"
	"blogspot/repository"
	"database/sql"
	"fmt"
)

type blogFlow struct {
	repoBlog     repository.IBlogRepository
	repoTag      repository.ITagRepository
	repoCategory repository.ICategoryRepository
}

func NewBlogFlow(blogRepo repository.IBlogRepository, repoTag repository.ITagRepository, repoCategory repository.ICategoryRepository) *blogFlow {
	return &blogFlow{
		repoBlog:     blogRepo,
		repoTag:      repoTag,
		repoCategory: repoCategory,
	}
}

func (f *blogFlow) GetBlog(blogFilter entity.BlogFilter) (blog entity.BlogResponse, err error) {
	blog, err = f.repoBlog.GetBlog(blogFilter)
	if err != nil {
		return
	}

	tagFilter := entity.TagFilter{
		BlogID: blog.ID,
	}

	categoryFilter := entity.CategoryFilter{
		BlogID: blog.ID,
	}

	blog.Tag, err = f.repoTag.GetAll(tagFilter)
	if err != nil {
		return
	}

	blog.Category, err = f.repoCategory.GetAll(categoryFilter)
	if err != nil {
		return
	}

	return blog, err
}

func (f *blogFlow) GetByUser(blogFilter entity.BlogFilter) (blogs []entity.BlogResponse, err error) {
	var blogsResponse []entity.BlogResponse

	blogs, err = f.repoBlog.GetByUser(blogFilter)
	if err != nil {
		fmt.Println("err 54", err)
		return
	}

	for _, blog := range blogs {
		tagFilter := entity.TagFilter{
			BlogID: blog.ID,
		}

		categoryFilter := entity.CategoryFilter{
			BlogID: blog.ID,
		}

		blog.Tag, err = f.repoTag.GetAll(tagFilter)
		if err != nil {
			return
		}

		fmt.Println(blog.Tag)

		blog.Category, err = f.repoCategory.GetAll(categoryFilter)
		if err != nil {
			return
		}

		blogsResponse = append(blogsResponse, blog)
	}

	return blogsResponse, err
}

func (f *blogFlow) GetHidden(blogFilter entity.BlogFilter) (blogs []entity.BlogResponse, err error) {
	var blogsResponse []entity.BlogResponse

	blogs, err = f.repoBlog.GetHidden(blogFilter)
	if err != nil {
		fmt.Println("err 89", err)
		return
	}

	for _, blog := range blogs {
		tagFilter := entity.TagFilter{
			BlogID: blog.ID,
		}

		categoryFilter := entity.CategoryFilter{
			BlogID: blog.ID,
		}

		blog.Tag, err = f.repoTag.GetAll(tagFilter)
		if err != nil {
			fmt.Println("err 104", err)
			return
		}

		blog.Category, err = f.repoCategory.GetAll(categoryFilter)
		if err != nil {
			fmt.Println("err 110", err)
			return
		}

		blogsResponse = append(blogsResponse, blog)
	}

	return blogsResponse, err
}

func (f *blogFlow) Insert(blog entity.Blog) (err error) {
	var tx *sql.Tx

	tx, err = f.repoBlog.BlogBeginTransaction()
	if err != nil {
		return err
	}

	lastInsertID, err := f.repoBlog.Insert(tx, blog)
	if err != nil {
		return
	}

	for _, tag := range blog.Tag {
		err = f.repoBlog.InsertTagDetail(tx, lastInsertID, tag)
		if err != nil {
			return
		}
	}

	for _, category := range blog.Category {
		err = f.repoBlog.InsertCategoryDetail(tx, lastInsertID, category)
		if err != nil {
			return
		}
	}

	if err = tx.Commit(); err != nil {
		return
	}

	return
}

func (f *blogFlow) Update(blog entity.Blog) (err error) {
	err = f.repoBlog.Update(blog)
	return
}

func (f *blogFlow) Hidden(blogFilter entity.BlogFilter) (err error) {
	err = f.repoBlog.Hidden(blogFilter)
	fmt.Println(err)
	return
}

func (f *blogFlow) Delete(blogID int) (err error) {
	var tx *sql.Tx

	tx, err = f.repoBlog.BlogBeginTransaction()
	if err != nil {
		fmt.Println("err 166", err)
		return err
	}

	err = f.repoBlog.DeleteTagDetail(tx, blogID)
	if err != nil {
		fmt.Println("err 172", err)
		return err
	}

	err = f.repoBlog.DeleteCategoryDetail(tx, blogID)
	if err != nil {
		fmt.Println("err 178", err)
		return err
	}

	err = f.repoBlog.Delete(tx, blogID)
	if err != nil {
		fmt.Println("err 184", err)
		return err
	}

	if err = tx.Commit(); err != nil {
		return
	}

	return
}
