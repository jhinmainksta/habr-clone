package repository

import "gorm.io/gorm"

type Repository struct {
	HabrClone
}

type HabrClone interface {
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		HabrClone: NewHabrClonePG(db),
	}
}
