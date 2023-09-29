package flow

import (
	"blogspot/entity"
	"blogspot/repository"
)

type categoryFlow struct {
	repo repository.ICategoryRepository
}

func NewCategoryFlow(categoryRepo repository.ICategoryRepository) *categoryFlow {
	return &categoryFlow{
		repo: categoryRepo,
	}
}

func (f *categoryFlow) Insert(category entity.Category) (err error) {
	err = f.repo.Insert(category)
	return
}
