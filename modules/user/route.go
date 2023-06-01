package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouterUser struct {
	UserRequestHandler RequestHandler
}

func NewRouter(db *gorm.DB) RouterUser {
	return RouterUser{
		UserRequestHandler: NewRequestHandler(db),
	}
}

func (r *RouterUser) Handle(router *gin.Engine) {
	basepath := "/user"
	user := router.Group(basepath)
	user.GET("", r.UserRequestHandler.RequestGetAllUser)
	user.GET("/:id", r.UserRequestHandler.RequestGetUserById)
	user.GET("/search", r.UserRequestHandler.RequestFilterUser)
	user.POST("",
		r.UserRequestHandler.RequestCreateUser,
	)
	user.PUT("/:id",
		r.UserRequestHandler.RequestUpdateUser,
	)
	user.DELETE("/:id",
		r.UserRequestHandler.RequestDeleteUser,
	)

}
