package repository

import (
	"gorm.io/gorm"
	"sesi1-crud/entity"
)

type RepositoryUser struct {
	db *gorm.DB
}

type RepositoryUserInterface interface {
	GetAllUsers() ([]entity.Users, error)
	GetUserByID(userId uint) (entity.Users, error)
	GetUserByName(username string) (entity.Users, error)
	CreateUser(in *entity.Users) (entity.Users, error)
	UpdateUser(in *entity.Users) (entity.Users, error)
	DeleteUser(userId uint) error
}

func NewRepoUser(db *gorm.DB) RepositoryUser {
	return RepositoryUser{
		db: db,
	}
}

func (ur *RepositoryUser) GetAllUsers() ([]entity.Users, error) {
	var users []entity.Users
	err := ur.db.Find(&users).Error
	return users, err
}

func (ur *RepositoryUser) GetUserByID(userId uint) (entity.Users, error) {
	var user entity.Users
	err := ur.db.Model(&user).Where("id = ?", userId).First(&user).Error
	return user, err
}

func (ur *RepositoryUser) GetUserByName(username string) ([]entity.Users, error) {
	var user []entity.Users
	err := ur.db.Where("name like ?", "%"+username+"%").Find(&user).Error
	return user, err
}

func (ur *RepositoryUser) CreateUser(in *entity.Users) (entity.Users, error) {
	err := ur.db.Model(&entity.Users{}).Create(in).Error
	return *in, err
}

func (ur *RepositoryUser) UpdateUser(in *entity.Users, userId uint) (entity.Users, error) {
	err := ur.db.Exec("UPDATE users SET name=?, email=?, password=? WHERE id=?",
		in.Name,
		in.Email,
		in.Password,
		userId,
	).Error
	return *in, err
}

func (ur *RepositoryUser) DeleteUser(userId uint) error {
	var users entity.Users
	users.ID = userId
	err := ur.db.First(&users).Where("id = ?", userId).Delete(&users).Error
	return err
}
