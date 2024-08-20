package Repository

import (
	"Golang/Models"
	"Golang/utils/Database"
	"gorm.io/gorm"
)

type (
	TaskRepository interface {
		PaginateWith(page, limit int, relations ...string) ([]Models.Task, int, error)
		Create(task *Models.Task) (*Models.Task, error)
		Update(task *Models.Task) (*Models.Task, error)
		Delete(id int) error
		FindById(id int) (Models.Task, error)
		FindByIdWith(id int, relations ...string) (Models.Task, error)
	}

	TaskRepositoryImpl struct {
		db *gorm.DB
		Paginator
	}
)

func (t TaskRepositoryImpl) FindByIdWith(id int, relations ...string) (Models.Task, error) {
	var task Models.Task
	result := t.db.Model(&Models.Task{}).Where("id = ?", id)
	Database.LoadRelations(result, relations...)
	result.Find(&task)
	return task, result.Error
}

func (t TaskRepositoryImpl) FindById(id int) (Models.Task, error) {
	var task Models.Task
	result := t.db.Model(&Models.Task{}).Where("id = ?", id).Find(&task)
	return task, result.Error
}

func (t TaskRepositoryImpl) PaginateWith(page, limit int, relations ...string) ([]Models.Task, int, error) {
	var tasks []Models.Task
	var totalRecords int64

	t.db.Model(&Models.Task{}).Count(&totalRecords)
	result := t.paginate(page, limit)
	Database.LoadRelations(result, relations...)
	result.Find(&tasks)

	return tasks, int(totalRecords), result.Error
}

func (t TaskRepositoryImpl) Create(task *Models.Task) (*Models.Task, error) {
	result := t.db.Create(task)
	return task, result.Error
}

func (t TaskRepositoryImpl) Update(task *Models.Task) (*Models.Task, error) {
	result := t.db.Save(task)
	return task, result.Error
}

func (t TaskRepositoryImpl) Delete(id int) error {
	result := t.db.Where("id = ?", id).Delete(&Models.Task{})
	return result.Error
}

func NewTaskRepository(db *gorm.DB, paginator *Paginator) TaskRepository {
	return &TaskRepositoryImpl{db: db, Paginator: *paginator}
}
