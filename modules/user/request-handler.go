package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sesi1-crud/payload/response"
	"sesi1-crud/repository"
	"strconv"
)

type RequestHandler struct {
	ctrl ControllerInterface
}

type RequestHandlerInterface interface {
	RequestGetAllUser(c *gin.Context)
	RequestGetUserByID(c *gin.Context)
	RequestFilterUser(c *gin.Context)
	RequestCreateUser(c *gin.Context)
	RequestUpdateUser(c *gin.Context)
	RequestDeleteUser(c *gin.Context)
}

func NewRequestHandler(db *gorm.DB) RequestHandler {
	return RequestHandler{
		ctrl: &Controller{
			UserUsecase: &Usecase{
				userRepo: repository.NewRepoUser(db),
			},
		},
	}
}
func (rh *RequestHandler) RequestGetAllUser(c *gin.Context) {
	res, err := rh.ctrl.GetUserAll()
	if err != nil {
		c.JSON(res.Status, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (rh *RequestHandler) RequestGetUserById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.HandleFailedResponse("ID Not Valid", 400))
		return
	}
	res, err := rh.ctrl.GetUserById(int(id))
	if err != nil {
		c.JSON(res.Status, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (rh *RequestHandler) RequestFilterUser(c *gin.Context) {
	inputName := c.Query("name")
	if inputName == "" {
		res, err := rh.ctrl.GetUserAll()
		if err != nil {
			c.JSON(res.Status, res)
			return
		}
		c.JSON(http.StatusOK, res)
		return
	}
	res, err := rh.ctrl.GetUserByName(inputName)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (rh *RequestHandler) RequestCreateUser(c *gin.Context) {
	request := UserDTO{}
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.HandleFailedResponse("request not valid", 400))
		return
	}
	res, err := rh.ctrl.CreateUser(request)
	if err != nil {
		c.JSON(res.Status, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (rh *RequestHandler) RequestUpdateUser(c *gin.Context) {
	userid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.HandleFailedResponse("id not valid", 400))
		return
	}
	request := UserDTO{}
	err = c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.HandleFailedResponse("request not valid", 400))
		return
	}
	res, err := rh.ctrl.UpdateUser(request, int(userid))
	if err != nil {
		c.JSON(res.Status, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (rh *RequestHandler) RequestDeleteUser(c *gin.Context) {
	userid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.HandleFailedResponse("id not valid", 400))
		return
	}
	res, err := rh.ctrl.DeleteUser(int(userid))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}
