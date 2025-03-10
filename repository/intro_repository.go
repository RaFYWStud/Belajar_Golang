package repository

import (
	"golang-tutorial/contract"
	"golang-tutorial/entity"

	"gorm.io/gorm"
)

type introRepo struct {
	db *gorm.DB
}

func implIntroRepository(db *gorm.DB) contract.IntroRepository {
	return &introRepo{
		db: db,
	}
}

func (r *introRepo) GetIntro(id int) (*entity.Intro, error) {
	var intro entity.Intro
	err := r.db.Table("intro").Where("id = ?", id).First(&intro).Error
	if err != nil {
		return nil, err
	}
	return &intro, err
}

func (r *introRepo) CreateIntro(intro *entity.Intro) error {
	return r.db.Table("intro").Create(intro).Error
}
