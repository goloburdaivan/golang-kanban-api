package Repository

import (
	"Golang/Models"
	"gorm.io/gorm"
)

type (
	ProjectRepository interface {
		GetAll() []Models.Project
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
	db := p.db.Find(&projects)
	if db.Error != nil {
		panic(db.Error)
	}

	return projects
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
