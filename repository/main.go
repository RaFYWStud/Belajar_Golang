package repository

import (
	"golang-tutorial/contract"

	"gorm.io/gorm"
)

func New(db *gorm.DB) *contract.Repository {
	return &contract.Repository{
		// Code here
		// Example:
		// Example: implExampleRepository(db),
		Intro: implIntroRepository(db),
		ToDo: implToDoRepository(db),
		Account: implAccountRepository(db),
	}
}
