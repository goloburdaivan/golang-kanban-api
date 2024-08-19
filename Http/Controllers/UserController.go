package Controllers

import (
	"Golang/DTO"
	"Golang/Http/Requests"
	"Golang/Repository"
	"Golang/Services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type (
	UserController interface {
		Login(c *gin.Context)
		Register(c *gin.Context)
		ConfirmEmail(c *gin.Context)
	}

	UserControllerImpl struct {
		repository          Repository.UserRepository
		confirmationService Services.EmailConfirmationService
		userService         Services.UserService
	}
)

func (u UserControllerImpl) Login(c *gin.Context) {
	var request Requests.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, token, err := u.userService.AuthenticateUser(request.Login, request.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
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

	user, token, err := u.userService.RegisterUser(request.Username, request.Login, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	go func() {
		err := u.confirmationService.Send(&DTO.ConfirmationDTO{
			User:  *user,
			Token: token,
		})
		if err != nil {
			log.Println(err)
		}
	}()

	c.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}

func (u UserControllerImpl) ConfirmEmail(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token required"})
		return
	}

	err := u.userService.ConfirmUserEmail(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Thank you for confirming your email!"})
}

func NewUserController(
	repository Repository.UserRepository,
	confirmationService Services.EmailConfirmationService,
	userService Services.UserService,
) UserController {
	return &UserControllerImpl{
		repository:          repository,
		confirmationService: confirmationService,
		userService:         userService,
	}
}
