package user

import (
	"sesi1-crud/entity"
	"sesi1-crud/repository"
)

type Usecase struct {
	userRepo repository.RepositoryUser
}

type UsecaseInterface interface {
	GetUserAll() ([]entity.Users, error)
	GetUserByID(userId int) (entity.Users, error)
	GetUserByName(username string) ([]entity.Users, error)
	CreateUser(in UserDTO) (entity.Users, error)
	UpdateUser(in UserDTO, userId uint) (entity.Users, error)
	DeleteUser(userId int) error
}

func (u *Usecase) GetUserAll() ([]entity.Users, error) {
	users, err := u.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *Usecase) GetUserByID(userId int) (entity.Users, error) {
	get, err := u.userRepo.GetUserByID(uint(userId))
	if err != nil {
		return entity.Users{}, err
	}
	return get, nil
}

func (u *Usecase) GetUserByName(username string) ([]entity.Users, error) {
	get, err := u.userRepo.GetUserByName(username)
	if err != nil {
		return nil, err
	}
	return get, nil
}

func (u *Usecase) CreateUser(in UserDTO) (entity.Users, error) {
	userSave := entity.Users{
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
	}
	get, err := u.userRepo.CreateUser(&userSave)
	if err != nil {
		return entity.Users{}, err
	}
	return get, nil
}

func (u *Usecase) UpdateUser(in UserDTO, userId uint) (entity.Users, error) {
	_, err := u.userRepo.GetUserByID(uint(userId))
	if err != nil {
		return entity.Users{}, err
	}
	userSave := entity.Users{
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
	}
	get, err := u.userRepo.UpdateUser(&userSave, userId)

	if err != nil {
		return entity.Users{}, err
	}
	return get, nil
}

func (u *Usecase) DeleteUser(userId int) error {
	err := u.userRepo.DeleteUser(uint(userId))

	if err != nil {
		return err
	}
	return nil
}
