package setup

import (
	"Golang/Models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

func InitDatabase() *gorm.DB {
	port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))

	var (
		host     = os.Getenv("DATABASE_HOST")
		user     = os.Getenv("DATABASE_USERNAME")
		password = os.Getenv("DATABASE_PASSWORD")
		dbName   = os.Getenv("DATABASE_NAME")
	)

	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", user, password, host, port, dbName)

	db, err := gorm.Open(mysql.Open(sqlInfo), &gorm.Config{})
	if err != nil {
		log.Print("Failed to connect to database")
		log.Fatal(err)
	}

	return db
}

func MigrateDatabase(db *gorm.DB) {
	err := db.AutoMigrate(
		&Models.User{},
		&Models.Project{},
		&Models.Task{},
		&Models.Column{},
		&Models.TaskComment{},
	)

	if err != nil {
		log.Print("Failed to migrate users table")
		log.Fatal(err)
	}
}
