package Database

import "gorm.io/gorm"

func LoadRelations(db *gorm.DB, relations ...string) {
	for _, relation := range relations {
		db.Preload(relation)
	}
}
