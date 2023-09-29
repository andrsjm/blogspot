package flow

import (
	"blogspot/entity"
	"blogspot/repository"
)

type tagFlow struct {
	repo repository.ITagRepository
}

func NewTagFlow(tagRepo repository.ITagRepository) *tagFlow {
	return &tagFlow{
		repo: tagRepo,
	}
}

func (f *tagFlow) Insert(tag entity.Tag) (err error) {
	err = f.repo.Insert(tag)
	return
}
