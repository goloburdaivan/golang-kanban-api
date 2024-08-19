package Repository

import (
	"Golang/Models"
	"gorm.io/gorm"
)

type (
	ProjectRepository interface {
		GetAll() []Models.Project
		Paginate(page int, limit int) ([]Models.Project, int, error)
		Find(id int) *Models.Project
		Create(project *Models.Project) *Models.Project
		Update(project *Models.Project) *Models.Project
		Delete(id int) *Models.Project
	}

	ProjectRepositoryImpl struct {
		db *gorm.DB
	}
)

func (p ProjectRepositoryImpl) GetAll() []Models.Project {
	var projects []Models.Project
	db := p.db.Preload("User").Find(&projects)
	if db.Error != nil {
		panic(db.Error)
	}

	return projects
}

func (p ProjectRepositoryImpl) Paginate(page int, limit int) ([]Models.Project, int, error) {
	var projects []Models.Project
	var count int64
	offset := (page - 1) * limit
	db := p.db.Preload("User").Offset(offset).Limit(limit).Find(&projects)
	p.db.Model(&Models.Project{}).Count(&count)
	return projects, int(count), db.Error
}

func (p ProjectRepositoryImpl) Find(id int) *Models.Project {
	var project Models.Project
	db := p.db.Where("id = ?", id).Find(&project)
	if db.Error != nil {
		panic(db.Error)
	}

	return &project
}

func (p ProjectRepositoryImpl) Create(project *Models.Project) *Models.Project {
	p.db.Create(project)
	return project
}

func (p ProjectRepositoryImpl) Update(project *Models.Project) *Models.Project {
	p.db.Save(project)
	return project
}

func (p ProjectRepositoryImpl) Delete(id int) *Models.Project {
	project := p.Find(id)
	p.db.Delete(project)
	return project
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &ProjectRepositoryImpl{db: db}
}
