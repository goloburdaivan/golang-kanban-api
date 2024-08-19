package Services

import "Golang/Repository"

type (
	TaskService interface {
		Move(id, columnId int)
	}
	TaskServiceImpl struct {
		repository Repository.TaskRepository
	}
)

func (t TaskServiceImpl) Move(id, columnId int) {
	task, _ := t.repository.FindById(id)
	task.ColumnID = columnId
	t.repository.Update(&task)
}

func NewTaskService(repository Repository.TaskRepository) TaskService {
	return &TaskServiceImpl{repository: repository}
}
