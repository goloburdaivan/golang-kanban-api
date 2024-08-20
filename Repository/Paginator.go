package Repository

import "gorm.io/gorm"

type Paginator struct {
	db *gorm.DB
}

func (p Paginator) paginate(page, limit int) *gorm.DB {
	return p.db.Offset((page - 1) * limit).Limit(limit)
}

func NewPaginator(db *gorm.DB) *Paginator {
	return &Paginator{db: db}
}
