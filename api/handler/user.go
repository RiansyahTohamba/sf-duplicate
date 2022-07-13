package handler

import (
	"errors"
	"fmt"
	"net/http"
	"sf-duplicate/repository"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	usrRepo *repository.UserRepository
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// TODO: extend later
type RegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewUserHandler(usrRepo *repository.UserRepository) *UserHandler {
	return &UserHandler{usrRepo}
}

func (usr *UserHandler) Login(ctx *gin.Context) {
	var loginReq LoginRequest
	var verr validator.ValidationErrors

	err := ctx.BindJSON(&loginReq)

	if err != nil {
		if errors.As(err, &verr) {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"errors": verr},
			)
		}
	}

	user, err := usr.usrRepo.Login(loginReq.Email, loginReq.Password)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
	} else {
		session := sessions.Default(ctx)
		userId := fmt.Sprintf("user:%d", user.ID)
		session.Set("user_id", userId)
		session.Save()

		ctx.JSON(http.StatusOK, gin.H{
			"message": "User Sign-in successfully",
		})

	}

}

func (usr *UserHandler) Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success logout",
	})
}
