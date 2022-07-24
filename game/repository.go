package game

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Game, error)
	FindById(ID int) (Game, error)
	Create(gim Game) (Game, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Game, error) {
	var gims []Game
	err := r.db.Find(&gims).Error
	return gims, err
}

func (r *repository) FindById(ID int) (Game, error) {
	var gim Game
	err := r.db.First(&gim, ID).Error
	return gim, err
}

func (r *repository) Create(gim Game) (Game, error) {
	err := r.db.Create(&gim).Error
	return gim, err
}
