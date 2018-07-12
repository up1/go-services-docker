package repository

import (
	"model"
)

type MyRepository interface {
	Inquiry() (*model.Item, error)
}

type mockRepository struct{}

func NewMockRepository() MyRepository {
	return &mockRepository{}
}

func (mr *mockRepository) Inquiry() (*model.Item, error) {
	return nil, nil
}
