package http

import (
	"go-starter-clean/pkg/entity"
	"go-starter-clean/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	log logger.LogInfoFormat
	usc entity.UserService
}

func NewUserHandler(log logger.LogInfoFormat, usc entity.UserService) *userHandler {
	return &userHandler{log: log, usc: usc}
}

func (u *userHandler) GetAll(ctx *gin.Context) {
	users, err := u.usc.FindAll()
	if len(users) == 0 || err != nil {
		ctx.Status(http.StatusNoContent)
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (u *userHandler) Store(ctx *gin.Context) {
	var user entity.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if errdb := u.usc.Store(&user); errdb != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errdb})
		return
	}
	ctx.Status(http.StatusCreated)

}
