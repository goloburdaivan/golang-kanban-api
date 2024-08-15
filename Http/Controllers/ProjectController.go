package Controllers

import (
	"Golang/DTO"
	"Golang/Models"
	"Golang/Repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type (
	ProjectController interface {
		Index(ctx *gin.Context)
		Show(ctx *gin.Context)
		Create(ctx *gin.Context)
		Update(ctx *gin.Context)
		Destroy(ctx *gin.Context)
	}

	ProjectControllerImpl struct {
		repository Repository.ProjectRepository
	}
)

func (p ProjectControllerImpl) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"projects": p.repository.GetAll(),
	})
}

func (p ProjectControllerImpl) Show(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	project := p.repository.Find(id)
	ctx.JSON(http.StatusOK, gin.H{
		"project": project,
	})
}

func (p ProjectControllerImpl) Create(ctx *gin.Context) {
	var project Models.Project
	err := ctx.ShouldBindJSON(&project)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	p.repository.Create(&project)

	ctx.JSON(http.StatusOK, gin.H{
		"project": project,
	})
}

func (p ProjectControllerImpl) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var projectDTO DTO.ProjectUpdateDTO
	err := ctx.ShouldBindJSON(&projectDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	project := p.repository.Find(id)
	p.repository.Update(project)

	ctx.JSON(http.StatusOK, gin.H{
		"project": project,
	})
}

func (p ProjectControllerImpl) Destroy(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	project := p.repository.Find(id)
	p.repository.Delete(id)
	ctx.JSON(http.StatusOK, gin.H{
		"project": project,
	})
}

func NewProjectController(repository Repository.ProjectRepository) ProjectController {
	return &ProjectControllerImpl{repository: repository}
}
