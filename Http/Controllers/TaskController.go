package Controllers

import (
	"Golang/Http/Requests"
	"Golang/Models"
	"Golang/Repository"
	"Golang/Services"
	"Golang/utils"
	"Golang/utils/Meta"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type (
	TaskController interface {
		Index(c *gin.Context)
		Create(c *gin.Context)
		Show(c *gin.Context)
		Update(c *gin.Context)
		Destroy(c *gin.Context)
		Move(c *gin.Context)
	}

	TaskControllerImpl struct {
		taskService Services.TaskService
		repository  Repository.TaskRepository
	}
)

func (t TaskControllerImpl) Index(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	tasks, totalRecords, err := t.repository.PaginateWith(page, limit, "Column.Project")

	if err != nil {
		utils.InternalServerErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"paginated": tasks,
		"pagination": &Meta.PaginationMeta{
			TotalRecords:   totalRecords,
			TotalPages:     (totalRecords + limit - 1) / limit,
			CurrentPage:    page,
			RecordsPerPage: limit,
		},
	})
}

func (t TaskControllerImpl) Create(c *gin.Context) {
	var task Models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		utils.InternalServerErrorResponse(c, err)
		return
	}

	_, err := t.repository.Create(&task)
	if err != nil {
		utils.InternalServerErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

func (t TaskControllerImpl) Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	task, err := t.repository.FindByIdWith(id, "Column.Project", "Assignee", "Creator")
	if err != nil {
		utils.InternalServerErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

func (t TaskControllerImpl) Update(c *gin.Context) {
	var task Models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		utils.InternalServerErrorResponse(c, err)
		return
	}

	_, err := t.repository.Update(&task)
	if err != nil {
		utils.InternalServerErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

func (t TaskControllerImpl) Destroy(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := t.repository.Delete(id)
	if err != nil {
		utils.InternalServerErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (t TaskControllerImpl) Move(c *gin.Context) {
	var request Requests.MoveTaskRequest
	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.InternalServerErrorResponse(c, err)
		return
	}

	t.taskService.Move(id, request.ColumnID)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Task moved to another column",
	})
}

func NewTaskController(
	taskService Services.TaskService,
	repository Repository.TaskRepository,
) TaskController {
	return &TaskControllerImpl{taskService: taskService, repository: repository}
}
