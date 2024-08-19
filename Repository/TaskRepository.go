package Repository

import (
	"Golang/Models"
	"gorm.io/gorm"
)

type (
	TaskRepository interface {
		Paginate(page, limit int) ([]Models.Task, int, error)
		Create(task *Models.Task) (*Models.Task, error)
		Update(task *Models.Task) (*Models.Task, error)
		Delete(id int) error
		FindById(id int) (Models.Task, error)
	}

	TaskRepositoryImpl struct {
		db *gorm.DB
	}
)

func (t TaskRepositoryImpl) FindById(id int) (Models.Task, error) {
	var task Models.Task
	result := t.db.Model(&Models.Task{}).Where("id = ?", id).Find(&task)
	return task, result.Error
}

func (t TaskRepositoryImpl) Paginate(page, limit int) ([]Models.Task, int, error) {
	var tasks []Models.Task
	var totalRecords int64
	t.db.Model(&Models.Task{}).Count(&totalRecords)
	err := t.db.Limit(limit).Offset((page - 1) * limit).Find(&tasks).Error
	return tasks, int(totalRecords), err
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

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &TaskRepositoryImpl{db: db}
}
