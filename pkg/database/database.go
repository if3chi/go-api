package database

import (
	"github.com/if3chi/go-api/pkg/kernel"
	"github.com/if3chi/go-api/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect(app *kernel.Application) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		app.Logger.Fatal(err.Error())
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Client{})

	return db
}
