package utils

import (
	"backend/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var m *migrate.Migrate

func InitMigrate() {
	gormDB := service.GetDB()
	db, _ := gormDB.DB()

	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, _ = migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"mysql",
		driver,
	)
}

func GetMigrate() *migrate.Migrate {
	return m
}
