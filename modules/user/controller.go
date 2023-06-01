package user

import (
	"net/http"
	"sesi1-crud/payload/response"
	"strconv"
)

type Controller struct {
	UserUsecase UsecaseInterface
}

type ControllerInterface interface {
	GetUserAll() (response.Response, error)
	GetUserByName(username string) (response.Response, error)
	GetUserById(userId int) (response.Response, error)
	CreateUser(in UserDTO) (response.Response, error)
	UpdateUser(in UserDTO, userId int) (response.Response, error)
	DeleteUser(userId int) (response.Response, error)
}

func (c *Controller) GetUserAll() (response.Response, error) {
	users, err := c.UserUsecase.GetUserAll()
	if err != nil {
		return response.HandleFailedResponse(err.Error(), http.StatusInternalServerError), err
	}
	return response.HandleSuccessResponse(users, "Success Get All User", 200), err
}

func (c *Controller) GetUserById(userId int) (response.Response, error) {
	user, err := c.UserUsecase.GetUserByID(userId)
	if err != nil {
		return response.HandleFailedResponse(err.Error(), http.StatusNotFound), err
	}
	return response.HandleSuccessResponse(user, "Success Get User By ID : "+strconv.Itoa(userId), 200), err
}

func (c *Controller) GetUserByName(username string) (response.Response, error) {
	user, err := c.UserUsecase.GetUserByName(username)
	if err != nil {
		return response.HandleFailedResponse(err.Error(), http.StatusNotFound), err
	}
	return response.HandleSuccessResponse(user, "Success Search User Contain Name : "+username, 200), err
}

func (c *Controller) CreateUser(in UserDTO) (response.Response, error) {
	_, err := c.UserUsecase.CreateUser(in)
	if err != nil {
		return response.HandleFailedResponse(err.Error(), http.StatusInternalServerError), err
	}
	return response.HandleSuccessResponse(nil, "Success Create User", 201), err
}

func (c *Controller) UpdateUser(in UserDTO, userId int) (response.Response, error) {
	_, err := c.UserUsecase.GetUserByID(userId)
	if err != nil {
		return response.HandleFailedResponse(err.Error(), http.StatusNotFound), err
	}
	user, err := c.UserUsecase.UpdateUser(in, uint(userId))
	if err != nil {
		return response.HandleFailedResponse(err.Error(), http.StatusInternalServerError), err
	}
	return response.HandleSuccessResponse(user, "Success Update User", 200), err
}

func (c *Controller) DeleteUser(userId int) (response.Response, error) {
	err := c.UserUsecase.DeleteUser(userId)
	if err != nil {
		return response.HandleFailedResponse(err.Error(), http.StatusInternalServerError), err
	}
	return response.HandleSuccessResponse(nil, "Success Delete User", 200), err
}
