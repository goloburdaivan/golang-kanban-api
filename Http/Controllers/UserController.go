package Controllers

import (
	"Golang/Http/Requests"
	"Golang/Repository"
	"Golang/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	UserController interface {
		Login(c *gin.Context)
		Register(c *gin.Context)
	}

	UserControllerImpl struct {
		repository Repository.UserRepository
	}
)

func (u UserControllerImpl) Login(c *gin.Context) {
	var request Requests.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := u.repository.FindByEmail(request.Login)
	if err != nil {
		utils.InternalServerErrorResponse(c, err)
		return
	}

	if !user.CheckPassword(request.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect password"})
		return
	}

	token, err := utils.GenerateAuthToken(user)
	if err != nil {
		utils.InternalServerErrorResponse(c, err)
		return
	}

	c.SetCookie("auth-token", token, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})
}

func (u UserControllerImpl) Register(c *gin.Context) {
	var request Requests.RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := u.repository.Create(request.Username, request.Login, request.Password)
	if err != nil {
		utils.InternalServerErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}

func NewUserController(repository Repository.UserRepository) UserController {
	return &UserControllerImpl{repository: repository}
}
