package repository

import "gorm.io/gorm"

type UserRepository struct {
	Db *gorm.DB
}

func (res *UserRepository) GetUser() {

}
